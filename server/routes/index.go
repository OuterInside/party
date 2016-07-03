package routes

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/OuterInside/party/server/entities"
	"github.com/OuterInside/party/server/models"
	"github.com/labstack/echo"
)

var (
	// set 2:10 (music play time)
	player  = models.CreatePlayer(2*time.Minute + 10*time.Second)
	clientM = &clientManager{
		Map: make(map[string]*entities.Client),
	}
)

// JSON map
type JSON map[string]interface{}

type clientManager struct {
	sync.Mutex

	Map map[string]*entities.Client
}

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

	random := make([]byte, 32)
	_, err = rand.Read(random)
	if err != nil {
		return
	}

	id := hex.EncodeToString(random)
	clientM.Lock()
	defer clientM.Unlock()
	clientM.Map[id] = &entities.Client{}

	return c.JSON(http.StatusOK, &entities.EntryResponse{
		ID:    id,
		Units: player.GetUnits(),
		Start: player.GetStartTime().Format(time.RFC3339),
	})
}

// 範囲内から出たイベント
func leave(c echo.Context) (err error) {
	clientM.Lock()
	defer clientM.Unlock()

	id := c.Param("id")
	log.Println("id:", id)

	if _, ok := clientM.Map[id]; ok {
		player.Stop()
		delete(clientM.Map, id)
		return c.JSON(http.StatusOK, &entities.LeaveResponse{
			Message: "ok",
		})
	}

	log.Printf("ID:%s not found!\n", id)
	return c.JSON(http.StatusBadRequest, &entities.LeaveResponse{
		Message: fmt.Sprintf("ID:%s not found!", id),
	})
}
