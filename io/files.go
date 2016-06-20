package io

import (
	"os"
)


// Determine whether the file exists.
func IsExist(name string) bool {
	_, err := os.Stat(name)
	return err == nil || os.IsExist(err)
}
