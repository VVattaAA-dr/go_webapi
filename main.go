// main.go
package main

import (
    "./src/controller"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // API namespace
    v1 := router.Group("/api/v1")
    {
        v1.GET("/tasks", controller.TasksGET)
    }

    router.Run(":8080")
}
