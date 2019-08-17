package test

import (
	"os"
	"fmt"
	"path"

	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
}

func (s Suite) GetFixture(filename string) string {
	filePath := os.Getenv("PWD")

	for {
		if path.Base(filePath) == "fapi" {
			break
		}

		filePath = path.Dir(filePath)
	}

	return fmt.Sprintf("%s/test/fixture/%s", filePath, filename)
}
