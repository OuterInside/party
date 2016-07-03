package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/OuterInside/party/server/entities"
	"github.com/OuterInside/party/server/models"
	"github.com/labstack/echo"
)

var (
	// set 2:10 (music play time)
	player = models.CreatePlayer(2*time.Minute + 10*time.Second)
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
	player.Play()

	return c.JSON(http.StatusOK, &entities.EntryResponse{
		ID:    "test",
		Units: player.GetUnits(),
		Start: player.GetStartTime().Format(time.RFC3339),
	})
}

// 範囲内から出たイベント
func leave(c echo.Context) (err error) {
	player.Stop()

	id := c.Param("id")
	log.Println("id:", id)
	return c.JSON(http.StatusOK, &entities.LeaveResponse{})
}
