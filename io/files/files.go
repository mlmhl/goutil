package io

import (
	"os"
)

// Create a new file
func Create(name string) (*os.File, error) {
	return os.Create(name)
}

// Determine whether the file exists.
func IsExist(name string) bool {
	_, err := os.Stat(name)
	return err == nil || os.IsExist(err)
}

func Remove(name string) error {
	return os.Remove(name)
}