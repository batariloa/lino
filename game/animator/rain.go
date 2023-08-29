package animator

import (
	"sync"
	"time"
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

func (a *RainAnimator) Animate(frameToBeDrawn *int) {

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		a.mu.Lock()
		*a.Timer++
		if *a.Timer > 8 {
			*a.Timer = 1
		}
		a.mu.Unlock()

		*frameToBeDrawn = a.chooseFrame()
	}
}

func (a *RainAnimator) chooseFrame() int {

	switch *a.Timer {

	case 1:
		return 380
	case 2:
		return 381
	case 3:
		return 382
	case 4:
		return 383
	case 5:
		return 384
	case 6:
		return 385
	case 7:
		return 386
	case 8:
		return 387
	default:
		return 0
	}
}
