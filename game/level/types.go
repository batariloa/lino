// level/types.go
package level

// Level represents a complete game level with all its components
type Level struct {
	ID         string     `json:"id"`
	Width      int        `json:"width"`
	Height     int        `json:"height"`
	Layers     [][]int    `json:"layers"`     // [0] is map layer, [1] is interaction layer
	Triggers   []Trigger  `json:"triggers"`   // Special areas that trigger events
	Properties Properties `json:"properties"` // Level-specific properties
}

// Properties holds level-specific state and settings
type Properties struct {
	LightEnabled bool                   `json:"lightEnabled"`
	StartingPosX int                    `json:"startingPosX"`
	StartingPosY int                    `json:"startingPosY"`
	Data         map[string]interface{} `json:"data,omitempty"` // Additional level data
}

// Position represents a 2D coordinate
type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Trigger represents a position-based event in the level
type Trigger struct {
	Position   Position               `json:"position"`             // Location of the trigger
	Type       string                 `json:"type"`                 // Type of trigger (e.g., "teleport")
	TargetRoom string                 `json:"targetRoom"`           // For teleport triggers
	Properties map[string]interface{} `json:"properties,omitempty"` // Additional trigger properties
}

// LayerType constants for different layer types
const (
	LayerMap         = 0 // Main map layer (walls, floors, etc)
	LayerInteraction = 1 // Interaction layer (switches, items, etc)
)

// TriggerType constants for different trigger types
const (
	TriggerTypeTeleport = "teleport"
	TriggerTypeDialog   = "dialog"
	TriggerTypeEvent    = "event"
)

// TileType constants for different tile types
const (
	TileEmpty    = 0
	TileWall     = 1
	TileDoor     = 2
	TileSwitch   = 3
	TileLightOn  = 67
	TileLightOff = 68
)

// State represents the current state of the level system
type State struct {
	CurrentLevel *Level
}

// Global state (if needed for compatibility with existing code)
var (
	CurrentLevel *Level
	LightOn      = true
)
