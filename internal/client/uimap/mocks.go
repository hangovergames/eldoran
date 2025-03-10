// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package uimap

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hangovergames/eldoria/internal/common/dtos"
	"github.com/stretchr/testify/mock"
)

// MockTileGrid mocks the TileGrid for testing purposes.
type MockTileGrid struct {
	mock.Mock
}

// Draw simulates drawing the tile grid.
func (m *MockTileGrid) Draw(screen *ebiten.Image, tileSizeX, tileSizeY int) {
	m.Called(screen, tileSizeX, tileSizeY)
}

// SetTile simulates setting a tile at a given position.
func (m *MockTileGrid) SetTile(x, y int, tileNames ...string) {
	m.Called(x, y, tileNames)
}

// GetTile simulates getting tile names at a given position.
func (m *MockTileGrid) GetTile(x, y int) ([]string, bool) {
	args := m.Called(x, y)
	return args.Get(0).([]string), args.Bool(1)
}

// Implement the LoadTileConfigDTOs method to satisfy the ITileGrid interface.
func (m *MockTileGrid) LoadTileConfigDTOs(tileConfigs []dtos.TileConfigDTO) {
	m.Called(tileConfigs)
}
