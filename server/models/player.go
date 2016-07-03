package models

import (
	"log"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"github.com/OuterInside/party/server/entities"
)

const playOffset = 0 * time.Second

// Player structure
type Player struct {
	startTime time.Time // 再生開始時刻
	startM    sync.Mutex

	units      int64         // 再生台数
	totalUnits int64         // 述べ再生台数
	duration   time.Duration // 楽曲再生時間[sec]

	parts  entities.Parts // パート毎のクライアント数
	partsM sync.Mutex
}

// CreatePlayer method
func CreatePlayer(duration time.Duration, partSize int) *Player {
	player := &Player{
		duration: duration,
		parts:    make([]entities.Part, partSize),
	}

	for i := range player.parts {
		player.parts[i] = entities.Part{
			ID: i,
		}
	}

	return player
}

func (player *Player) start() {
	t := <-time.After(player.duration)
	log.Println("player end:", t)

	// reset
	atomic.StoreInt64(&player.totalUnits, 0)
	atomic.StoreInt64(&player.units, 0)
}

// Play method
func (player *Player) Play() (partID int) {
	player.partsM.Lock()
	defer player.partsM.Unlock()

	// 1台目
	if atomic.LoadInt64(&player.totalUnits) == 0 {
		player.startM.Lock()
		defer player.startM.Unlock()
		player.startTime = time.Now().Add(playOffset)
		go player.start()

		atomic.StoreInt64(&player.units, 0)
	} else {
		// copy
		list := make(entities.Parts, len(player.parts))
		for i, v := range player.parts {
			list[i] = v
		}
		// sort
		sort.Sort(list)
		// log.Println("list:", list)
		partID = list[0].ID
	}

	// increment units
	atomic.AddInt64(&player.totalUnits, 1)
	atomic.AddInt64(&player.units, 1)

	log.Println("totalUnits:", atomic.LoadInt64(&player.totalUnits))
	log.Println("units:", atomic.LoadInt64(&player.units))

	// increment part
	atomic.AddInt64(&player.parts[partID].Count, 1)
	// log.Println("player.parts:", player.parts)

	return
}

// Stop method
func (player *Player) Stop(partID int) {
	atomic.AddInt64(&player.units, -1)

	// decrement part
	player.partsM.Lock()
	defer player.partsM.Unlock()
	atomic.AddInt64(&player.parts[partID].Count, -1)
	// log.Println("player.parts:", player.parts)

	log.Println("totalUnits:", atomic.LoadInt64(&player.totalUnits))
	log.Println("units:", atomic.LoadInt64(&player.units))
}

// GetStartTime method
func (player *Player) GetStartTime() time.Time {
	player.startM.Lock()
	defer player.startM.Unlock()
	return player.startTime
}

// GetUnits method
func (player *Player) GetUnits() int64 {
	return atomic.LoadInt64(&player.units)
}
