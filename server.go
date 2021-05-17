package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func main() {
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
    log.Printf("defaulting to port %s", port)
  }
  e := echo.New()
  e.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
  })
  e.GET("/health", func(c echo.Context) error {
    return c.JSON(http.StatusOK, map[string]interface{}{
      "status": "ok",
    })
  })
  e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
