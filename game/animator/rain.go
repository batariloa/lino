package animator

import (
	"sync"
	"time"
)

const (
	startRainFrame = 379
)

type RainAnimator struct {
	Timer *int
	mu    sync.Mutex
}

func NewRainAnimator() *RainAnimator {
	timer := 0
	return &RainAnimator{
		Timer: &timer,
	}
}

func (a *RainAnimator) StartTimer() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		a.mu.Lock()
		*a.Timer++
		if *a.Timer > 8 {
			*a.Timer = 1
		}
		a.mu.Unlock()
	}
}

func (a *RainAnimator) CurrentRainFrame() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return startRainFrame + *a.Timer
}
