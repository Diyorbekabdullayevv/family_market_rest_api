package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	dbconnect "practice_gin.com/internal/dbConnect"
	"practice_gin.com/internal/models"
)

func GetProducts(context *gin.Context) {

	id := context.Param("id")

	db, err := dbconnect.ConnectDB()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to CONNECT to database!"})
		fmt.Println("Error:", err)
	}
	defer db.Close()

	var product models.Products
	err = db.QueryRow(`SELECT product_name, product_type, description, category, brand, is_available FROM products WHERE product_id = ?`, id).
		Scan(&product.Name, &product.Type, &product.Description, &product.Category, &product.Brand, &product.IsAvailable)
	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusInternalServerError, gin.H{"Error": "Product not found!"})
			fmt.Println("Product not found:", err)
		}
		context.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to SEND a query to database!"})
		fmt.Println("Failed to SEND a query to database:", err)
	}
	context.JSON(http.StatusOK, gin.H{"Product": product})
	// context.JSON(http.StatusOK, gin.H{"Content": "My GET request response!"})
}

func PostProducts(context *gin.Context) {
	var products []models.Products
	err := context.ShouldBindBodyWithJSON(&products)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to DECODE request body!"})
		fmt.Println("Failed to DECODE request body:", err)
	}

	db, err := dbconnect.ConnectDB()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to CONNECT to database!"})
		fmt.Println("Failed to CONNECT to database:", err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT product_name, product_type, description, category, brand, is_available FROM products`)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to SEND query to database!"})
		fmt.Println("Failed to SEND query to database:", err)
	}
	defer rows.Close()

	var retrievedProductsList []models.Products
	for rows.Next() {
		var retrievedProduct models.Products
		err = rows.Scan(&retrievedProduct.Name, &retrievedProduct.Type, &retrievedProduct.Description, &retrievedProduct.Category, &retrievedProduct.Brand, &retrievedProduct.IsAvailable)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to RETRIEVE data from database!"})
			fmt.Println("Failed to RETRIEVE data from database:", err)
		}
		retrievedProductsList = append(retrievedProductsList, retrievedProduct)
	}

	for _, product := range products {

		for _, retrievedProduct := range retrievedProductsList {
			if retrievedProduct.Name == product.Name {
				context.JSON(http.StatusBadRequest, gin.H{"Error": "Product with this name already exists!"})
				return
			}
		}

		_, err = db.Exec(`INSERT INTO products(product_name, product_type, description, category, brand) VALUES(?,?,?,?,?)`,
			product.Name, product.Type, product.Description, product.Category, product.Brand)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to EXECUTE data to database!"})
			fmt.Println("Failed to EXECUTE data to database:", err)
		}
	}
	context.JSON(http.StatusCreated, gin.H{"Status": "Products have been successfully stored in the database!"})
	context.JSON(http.StatusCreated, gin.H{"Status": "Something new happened here!"})
}

func GetStructValues(model interface{}) []interface{} {
	modelValue := reflect.ValueOf(model)
	modelType := modelValue.Type()
	values := []interface{}{}
	for i := 0; i < modelType.NumField(); i++ {
		dbTag := modelType.Field(i).Tag.Get("db")
		if dbTag != "" && dbTag != "id,omitempty" {
			values = append(values, modelValue.Field(i).Interface())
		}
	}
	return values
}
