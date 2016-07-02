package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/OuterInside/party/server/entities"
	"github.com/labstack/echo"
)

// JSON map
type JSON map[string]interface{}

// New method
func New(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &JSON{
			"message": "PartyServer works!",
		})
	})

	e.POST("/enter", enter)
	e.GET("/enter", enter)

	e.POST("/leave/:id", leave)
	e.GET("/leave/:id", leave)
}

// 範囲内に入ったイベント
func enter(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, &entities.EntryResponse{
		ID:    "test",
		Units: 1,
		Start: time.Now().Add(10 * time.Second).Format(time.RFC3339),
	})
}

// 範囲内から出たイベント
func leave(c echo.Context) (err error) {
	id := c.Param("id")
	log.Println("id:", id)
	return c.JSON(http.StatusOK, &entities.LeaveResponse{})
}
