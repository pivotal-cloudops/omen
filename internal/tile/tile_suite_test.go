package tile

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tile Suite")
}
