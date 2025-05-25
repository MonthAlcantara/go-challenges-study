package main

import (
	// Imports the domain package, where the Product struct and NewProduct function are defined.
	// In Go, it's common to organize code into small, cohesive packages, making maintenance and reuse easier.
	// In Java, you would typically use packages and classes for this purpose.
	"github.com.br/monthalcantara/products-challenge/internal/domain"
	// Imports the Gin framework to create HTTP APIs.
	// Gin is a popular Go library because it's lightweight, fast, and has a simple syntax.
	// Using minimalist libraries is a common practice in Go, prioritizing clarity and efficiency.
	// In Java, you might use Spring Boot or Jakarta EE for similar functionality.
	"github.com/gin-gonic/gin"
)

func main() {
	// Creates a new Gin server instance.
	// In Go, it's common to initialize dependencies directly in the main function, keeping the code simple and straightforward.
	// In Java, you would typically use SpringApplication.run or similar to bootstrap your application.
	server := gin.Default()

	// Defines a GET route for "/products".
	// The Go standard is to use anonymous functions (closures) for handlers, making the code more concise.
	// In Java, you would define a method annotated with @GetMapping("/products") in a controller class.
	server.GET("/products", func(c *gin.Context) {
		// Creates a list of products using the NewProduct function.
		// In Go, structs and constructor functions are preferred over classes, following the principle of simplicity.
		// In Java, you would instantiate objects using the 'new' keyword and constructors.
		products := []domain.Product{
			*domain.NewProduct("Product 1", 10.0, 5),
			*domain.NewProduct("Product 2", 20.0, 10),
		}
		// Returns the list of products as a JSON response.
		// Returning JSON is standard in REST APIs, and Gin makes this easy with simple methods.
		// In Java, you would typically return a List<Product> from a controller method, and frameworks like Spring handle JSON serialization automatically.
		c.JSON(200, products)
	})

	// Starts the server on port 8080.
	// In Go, it's common to specify the port directly, avoiding complex configurations.
	// In Java, the port is usually set in application.properties or via annotations/configuration.
	server.Run(":8080")
	// Note: Go style values clean, direct, and easy-to-read code, avoiding unnecessary abstractions.
}
