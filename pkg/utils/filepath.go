package utils

import (
	"errors"
	"os"

	"github.com/rs/zerolog/log"
)

// EnsureDir: create the directory if doesn't exist
func EnsureDir(dir string) (string, error) {

	// TODO: get it lowercase

	err := os.MkdirAll(dir, 0755)
	if !errors.Is(err, os.ErrExist) && err != nil {
		log.Error().Err(err).Msgf("failed to create dir: %s", dir)
		return dir, err
	}

	// if _, err := os.Stat(dir); os.IsNotExist(err) {
	// 	log.Error().Err(err).Msgf("failed to create dir: %s", dir)
	//
	// 	err := os.MkdirAll(dir, 0755)
	//
	// 	if err != nil {
	// 		log.Debug().Err(err).Caller(0)
	// 	}
	//
	// 	return dir, err
	// }
	return dir, nil
}
