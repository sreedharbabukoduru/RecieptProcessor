package main

import (
    "github.com/gin-gonic/gin"
    "receipt-processor/internal/handlers"
)

func main() {
    router := gin.Default()

    // Define routes
    router.POST("/receipts/process", handlers.ProcessReceipt)
    router.GET("/receipts/:id/points", handlers.GetPoints)

    // Start the server
    router.Run(":8080")
}