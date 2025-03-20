package main

import (
	"encoding/json"
	"os"
)

type LayerJSON struct {
	Data   []int `json:"data"`
	Height int   `json:"height"`
	Width  int   `json:"width"`
}

type TileMapJSON struct {
	Layers []*LayerJSON `json:"layers"`
}

func LoadTileMapFromJSON(jsonPath string, tileMap *TileMapJSON) error {
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &tileMap)
	if err != nil {
		return err
	}

	return nil
}
