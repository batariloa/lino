// level/manager.go
package level

import (
	"fmt"
	"log"
)

type LevelManager struct {
	currentLevel *Level
	levels       map[string]*Level
}

func NewLevelManager() *LevelManager {
	return &LevelManager{
		levels: make(map[string]*Level),
	}
}

// LoadLevel loads a level by ID and sets it as current
func (lm *LevelManager) LoadLevel(id string) error {
	// Check if already loaded
	if level, exists := lm.levels[id]; exists {
		lm.currentLevel = level
		return nil
	}

	// Load new level
	level, err := loadLevelFromFile(fmt.Sprintf("resources/levels/%s.json", id))
	if err != nil {
		return fmt.Errorf("failed to load level %s: %w", id, err)
	}

	lm.levels[id] = level
	lm.currentLevel = level

	log.Printf("Loaded level: %s", id)
	return nil
}

// GetTileAt returns the map tile at the specified position
func (lm *LevelManager) GetTileAt(x, y int) int {
	if !lm.isValidPosition(x, y) {
		log.Println("Invalid tile position")
		return TileEmpty
	}

	index := y*lm.currentLevel.Width + x
	return lm.currentLevel.Layers[LayerMap][index]
}

// GetInteractionAt returns the interaction tile at the specified position
func (lm *LevelManager) GetInteractionAt(x, y int) int {
	if !lm.isValidPosition(x, y) {
		return TileEmpty
	}

	index := y*lm.currentLevel.Width + x
	return lm.currentLevel.Layers[LayerInteraction][index]
}

// GetTriggerAt returns any trigger at the specified position
func (lm *LevelManager) GetTriggerAt(x, y int) *Trigger {
	if lm.currentLevel == nil {
		return nil
	}

	for i := range lm.currentLevel.Triggers {
		trigger := &lm.currentLevel.Triggers[i]
		if trigger.Position.X == x && trigger.Position.Y == y {
			return trigger
		}
	}
	return nil
}

// ReplaceTile replaces a tile in the specified layer
func (lm *LevelManager) ReplaceTile(x, y int, layer int, newTile int) error {
	if !lm.isValidPosition(x, y) {
		return fmt.Errorf("invalid position: %d,%d", x, y)
	}
	if layer < 0 || layer >= len(lm.currentLevel.Layers) {
		return fmt.Errorf("invalid layer: %d", layer)
	}

	index := y*lm.currentLevel.Width + x
	lm.currentLevel.Layers[layer][index] = newTile
	return nil
}

// HandleTrigger processes a trigger at the given position
func (lm *LevelManager) HandleTrigger(x, y int) error {
	trigger := lm.GetTriggerAt(x, y)
	if trigger == nil {
		return nil
	}

	switch trigger.Type {
	case TriggerTypeTeleport:
		return lm.LoadLevel(trigger.TargetRoom)
	case TriggerTypeDialog:
		// Handle dialog trigger
		log.Printf("Dialog trigger activated at %d,%d", x, y)
		return nil
	case TriggerTypeEvent:
		// Handle event trigger
		log.Printf("Event trigger activated at %d,%d", x, y)
		return nil
	default:
		return fmt.Errorf("unknown trigger type: %s", trigger.Type)
	}
}

// GetCurrentLevel returns the current level
func (lm *LevelManager) GetCurrentLevel() *Level {
	return lm.currentLevel
}

// Helper method to check if a position is valid in the current level
func (lm *LevelManager) isValidPosition(x, y int) bool {
	if lm.currentLevel == nil {
		return false
	}

	if x < 0 || y < 0 || x >= lm.currentLevel.Width || y >= lm.currentLevel.Height {
		return false
	}

	index := y*lm.currentLevel.Width + x
	return index < len(lm.currentLevel.Layers[LayerMap])
}

// UpdateLevelProperty updates a property in the current level
func (lm *LevelManager) UpdateLevelProperty(key string, value interface{}) error {
	if lm.currentLevel == nil {
		return fmt.Errorf("no level loaded")
	}

	if lm.currentLevel.Properties.Data == nil {
		lm.currentLevel.Properties.Data = make(map[string]interface{})
	}

	lm.currentLevel.Properties.Data[key] = value
	return nil
}

// ToggleLight toggles the light state in the current level
func (lm *LevelManager) ToggleLight() error {
	if lm.currentLevel == nil {
		return fmt.Errorf("no level loaded")
	}

	lm.currentLevel.Properties.LightEnabled = !lm.currentLevel.Properties.LightEnabled
	return nil
}

// GetLightState returns the current light state
func (lm *LevelManager) GetLightState() bool {
	if lm.currentLevel == nil {
		return false
	}
	return lm.currentLevel.Properties.LightEnabled
}
