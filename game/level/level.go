package level

import (
	"encoding/json"
	"fmt"
	"os"
)

type LevelData struct {
	Map           [][]int        `json:"layers"`
	Interactables []Interactable `json:"interactables"`
	Triggers      []Trigger      `json:"triggers"`
}

type Interactable struct {
	X          int    `json:"x"`
	Y          int    `json:"y"`
	Type       string `json:"type"`
	TargetRoom string `json:"targetRoom"`
}

type Trigger struct {
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Action string `json:"action"`
}

type Level interface {
	GenerateMap()
	MoveRoom()
}

func LoadLevel(room string) error {
	filePath := fmt.Sprintf("./resources/levels/%s.json", room)

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file %s: %w", filePath, err)
	}
	defer file.Close()

	var levelData LevelData
	if err := json.NewDecoder(file).Decode(&levelData); err != nil {
		return fmt.Errorf("error decoding JSON from file %s: %w", filePath, err)
	}

	LevelMap = &levelData.Map
	return nil
}
