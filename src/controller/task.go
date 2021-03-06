// src/controller/task.go
package controller

import (
    // "database/sql"
    "fmt"
    "net/http"
    "time"

    "../model"
    "github.com/gin-gonic/gin"
)

// TasksGET returns list of tasks
func TasksGET(c *gin.Context) {
    db := model.DBConnect()
    result, err := db.Query("SELECT * FROM task ORDER BY id DESC")
    if err != nil {
        panic(err.Error())
    }

    tasks := []model.Task{}

    // iterate result
    for result.Next() {
        task := model.Task{}
        var id uint
        var createdAt, updatedAt time.Time
        var title string

        err = result.Scan(&id, &createdAt, &updatedAt, &title)
        if err != nil {
            panic(err.Error())
        }

        task.ID = id
        task.CreatedAt = createdAt
        task.UpdatedAt = updatedAt
        task.Title = title
        tasks = append(tasks, task)
    }
    fmt.Println(tasks)
    c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}
