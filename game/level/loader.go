// level/loader.go
package level

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// loadLevelFromFile loads and validates a level from a JSON file
func loadLevelFromFile(path string) (*Level, error) {
	// Get absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("invalid path: %w", err)
	}

	// Open and read file
	file, err := os.Open(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open level file: %w", err)
	}
	defer file.Close()

	// Decode JSON
	var level Level
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&level); err != nil {
		return nil, fmt.Errorf("failed to decode level data: %w", err)
	}

	// Validate level data
	if err := validateLevel(&level); err != nil {
		return nil, fmt.Errorf("invalid level data: %w", err)
	}

	return &level, nil
}

// validateLevel checks if the level data is valid
func validateLevel(level *Level) error {
	// Check basic level properties
	if level.ID == "" {
		return fmt.Errorf("level ID cannot be empty")
	}

	if level.Width <= 0 {
		return fmt.Errorf("invalid dimensions: width=%d, height=%d",
			level.Width, level.Height)
	}

	return nil
}
