package models

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
)

// Player structure
type Player struct {
	startTime time.Time // 再生開始時刻
	startM    sync.Mutex

	units      int64         // 再生台数
	totalUnits int64         // 述べ再生台数
	duration   time.Duration // 楽曲再生時間[sec]
}

// CreatePlayer method
func CreatePlayer(duration time.Duration) *Player {
	return &Player{duration: duration}
}

func (player *Player) start() {
	t := <-time.After(player.duration)
	log.Println("player end:", t)

	// reset
	atomic.StoreInt64(&player.totalUnits, 0)
	atomic.StoreInt64(&player.units, 0)
}

// Play method
func (player *Player) Play() {
	// 1台目
	if atomic.LoadInt64(&player.totalUnits) == 0 {
		player.startM.Lock()
		defer player.startM.Unlock()
		player.startTime = time.Now().Add(10 * time.Second)
		go player.start()

		atomic.StoreInt64(&player.units, 0)
	}

	// increment units
	atomic.AddInt64(&player.totalUnits, 1)
	atomic.AddInt64(&player.units, 1)

	log.Println("totalUnits:", atomic.LoadInt64(&player.totalUnits))
	log.Println("units:", atomic.LoadInt64(&player.units))
}

// Stop method
func (player *Player) Stop() {
	atomic.AddInt64(&player.units, -1)

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
