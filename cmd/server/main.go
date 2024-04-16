package main

import (
	// "net/http"
	// "example.com/server/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
    // Initialize routes
	router := gin.Default()

    // Start server
    router.Run(":8000")
}
