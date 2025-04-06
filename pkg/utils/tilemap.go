package utils

import (
	"encoding/json"
	"os"

	"github.com/majorbruteforce/clash2D/pkg/errorutils"
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
	errorutils.CheckFatal(err)

	tilemap := &TileMapJSON{}

	json.Unmarshal(data, tilemap)

	return tilemap
}
