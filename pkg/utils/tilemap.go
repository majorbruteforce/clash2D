package utils

import (
	"clash2D/pkg/exception"
	"encoding/json"
	"os"
)

type TileMapLayer struct {
	Data   []int `json:"data"`
	Height int   `json:"height"`
	Width  int   `json:"width"`
}

type TileMapJSON struct {
	Layers []*TileMapLayer `json:"layers"`
}

func ExtarctTileMapFromJSON(path string) *TileMapJSON {

	data, err := os.ReadFile(path)
	exception.CheckFatal(err)

	tilemap := &TileMapJSON{}

	json.Unmarshal(data, tilemap)

	return tilemap
}

// func parseJSON() {}
