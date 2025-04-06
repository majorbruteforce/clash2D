package config

import (
	"os"

	"github.com/majorbruteforce/clash2D/pkg/errorutils"
)

var RootDir string

func Init() {
	dir, err := os.Getwd()
	if err != nil {
		errorutils.CheckFatal(err)
	}

	RootDir = dir
}
