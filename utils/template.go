package utils

import (
	"os"
	"path/filepath"
)

// Creates absolute path to template file.
// Panics if it cannot get working directory
func GetTemplateFilepath(templatefile string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return filepath.Join(wd, "templates", templatefile)
}
