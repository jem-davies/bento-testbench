package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type input struct {
	Name string `json:"name"`
}

type output struct {
	Greeting string `json:"greeting"`
}

func main() {
	fmt.Printf("Booting test http server at address: %v\n", "localhost:8080")

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.POST("/greeting", greeting)
	router.Run("localhost:8080")
}

func greeting(c *gin.Context) {
	var in input

	if err := c.BindJSON(&in); err != nil {
		return
	}

	out := output{Greeting: "Hello " + in.Name}

	c.JSON(http.StatusAccepted, out)
}
