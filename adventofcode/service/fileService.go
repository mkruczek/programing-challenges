package service

import (
	"os"
)

func OpenFile(filePath string) (*os.File, error) {
	return os.Open(filePath)
}
