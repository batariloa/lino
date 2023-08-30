package animator

import (
	"math/rand"
	"sync"
	"time"
)

const (
	startRainFrame = 379
)

type RainAnimator struct {
	Timer      *int
	RainFrames map[int]int
	mu         sync.Mutex
}

func NewRainAnimator() *RainAnimator {
	timer := 0
	return &RainAnimator{
		Timer:      &timer,
		RainFrames: make(map[int]int),
	}
}

func (a *RainAnimator) StartTimer() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		a.mu.Lock()
		*a.Timer++
		if *a.Timer > 9 {
			*a.Timer = 1
		}

		// Update rain frames for each tile
		for tileID := range a.RainFrames {
			a.RainFrames[tileID] = (a.RainFrames[tileID] + 1) % 20
		}

		a.mu.Unlock()
	}
}

func (a *RainAnimator) CurrentRainFrame() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return startRainFrame + *a.Timer
}

func randomDuration() time.Duration {
	// Generate a random duration between 50ms and 200ms
	return time.Duration(rand.Intn(150)+50) * time.Millisecond
}
