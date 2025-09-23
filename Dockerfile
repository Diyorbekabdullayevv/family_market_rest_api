FROM golang

WORKDIR /docker_practice

COPY . .

EXPOSE 8080

CMD ["go", "run", "cmd/api/main.go"]