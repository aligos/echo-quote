package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	Excuse struct {
		Error  string `json:"error"`
		ID     string `json:"id"`
		Author string `json:"author"`
		Quote  string `json:"quote"`
	}
)

func getQuote(c echo.Context) error {
	requestedid := c.Param("id")
	db, err := sql.Open("mysql", "root:@/excuse")

	if err != nil {
		fmt.Println(err.Error())
		response := Excuse{ID: "", Error: "true", Author: "", Quote: ""}
		return c.JSON(http.StatusInternalServerError, response)
	}

	defer db.Close()

	var quote, id, author string
	if requestedid != "" {
		err = db.QueryRow("SELECT id, author, quote FROM excuses WHERE id = ?", requestedid).Scan(&id, &author, &quote)
	} else {
		err = db.QueryRow("SELECT id, author, quote FROM excuses ORDER BY RAND() LIMIT 1").Scan(&id, &author, &quote)
	}

	if err != nil {
		fmt.Println(err)
	}

	response := Excuse{ID: id, Error: "false", Author: author, Quote: quote}
	return c.JSON(http.StatusOK, response)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/quote", getQuote)
	e.GET("/id/:id", getQuote)

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
