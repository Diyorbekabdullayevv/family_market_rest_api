package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	dbconnect "practice_gin.com/internal/dbConnect"
	"practice_gin.com/internal/models"
)

func HomePage(ctx *gin.Context) {

	product := models.Product{
		Name:        "--------",
		Type:        "--------",
		Description: "--------",
	}
	ctx.HTML(http.StatusOK, "homepage.html", product)
}

func GetProductHomePage(ctx *gin.Context) {

	productName := ctx.PostForm("search_input")
	if productName != "" {
		db, err := dbconnect.ConnectDB()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to CONNECT to database!"})
			fmt.Println("Error:", err)
			return
		}
		defer db.Close()
		var product models.Product
		err = db.QueryRow(`SELECT name, type, description, cost FROM products WHERE name = ?`, productName).
			Scan(&product.Name, &product.Type, &product.Description, &product.Cost)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Product doesn`t exist with this name!"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to SEND a query to database!"})
			fmt.Println("Error:", err)
			return
		}
		ctx.HTML(http.StatusOK, "homepage.html", product)
	}
}

func PurchaseProducts(ctx *gin.Context) {

}

func GetProducts(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "get_products.html", nil)
}

func HandlerProductsForm(ctx *gin.Context) {
	product_id := ctx.PostForm("product_id")
	fmt.Println("id:", product_id)
	redirectUrl := fmt.Sprintf("/products/%v", product_id)
	fmt.Println(redirectUrl)
	ctx.Redirect(http.StatusFound, redirectUrl)
}

func GetProductByID(ctx *gin.Context) {

	id := ctx.Param("id")

	db, err := dbconnect.ConnectDB()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to CONNECT to database!"})
		fmt.Println("Error:", err)
	}
	defer db.Close()

	var product models.Product
	err = db.QueryRow(`SELECT name, type, description, is_available, cost FROM products WHERE id = ?`, id).
		Scan(&product.Name, &product.Type, &product.Description, &product.IsAvailable, &product.Cost)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Product not found!"})
			fmt.Println("Product not found:", err)
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to SEND a query to database!"})
		fmt.Println("Failed to SEND a query to database:", err)
	}

	ctx.HTML(http.StatusOK, "get_product.html", product)
}

func AddProducts(ctx *gin.Context) {
	var products []models.Product
	err := ctx.ShouldBindBodyWithJSON(&products)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to DECODE request body!"})
		fmt.Println("Failed to DECODE request body:", err)
	}

	db, err := dbconnect.ConnectDB()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to CONNECT to database!"})
		fmt.Println("Failed to CONNECT to database:", err)
	}
	defer db.Close()

	for _, product := range products {

		_, err = db.Exec(`
		INSERT INTO products(
			name,
			type,
			description,
			category_id,
			brand_id,
			is_available,
			unit_number,
			barcode,
			manufacture_date,
			receive_date,
			expiry_date,
			stock_keeping_date,
			cost,
			discount
		) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
			product.Name,
			product.Type,
			product.Description,
			product.CategoryID,
			product.BrandID,
			product.IsAvailable,
			product.UnitNumber,
			product.Barcode,
			product.ManufactureDate,
			product.ReceiveDate,
			product.ExpiryDate,
			product.StockKeepingDate,
			product.Cost,
			product.Discount,
		)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert product into database"})
			fmt.Println("Failed to insert product into database:", err)
			return
		}

	}
	ctx.JSON(http.StatusCreated, gin.H{"Status": "Products have been successfully stored in the database!"})
	ctx.JSON(http.StatusCreated, gin.H{"Status": "Something new happened here!"})
}

func AddCategories(ctx *gin.Context) {
	var categories []models.Category
	err := ctx.ShouldBindBodyWithJSON(&categories)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to DECODE request body!"})
		fmt.Println("Failed to DECODE request body:", err)
	}

	db, err := dbconnect.ConnectDB()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to CONNECT to database!"})
		fmt.Println("Failed to CONNECT to database:", err)
	}
	defer db.Close()

	for _, category := range categories {
		_, err = db.Exec(`INSERT INTO categories(name) VALUES(?)`, category.Name)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert product into database"})
			fmt.Println("Failed to insert product into database:", err)
			return
		}
	}
}

func AddBrands(ctx *gin.Context) {
	var brands []models.Brand
	err := ctx.ShouldBindBodyWithJSON(&brands)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to DECODE request body!"})
		fmt.Println("Failed to DECODE request body:", err)
	}

	db, err := dbconnect.ConnectDB()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to CONNECT to database!"})
		fmt.Println("Failed to CONNECT to database:", err)
	}
	defer db.Close()

	for _, brand := range brands {
		_, err = db.Exec(`INSERT INTO brands(name) VALUES(?)`, brand.Name)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert product into database"})
			fmt.Println("Failed to insert product into database:", err)
			return
		}
	}
}
