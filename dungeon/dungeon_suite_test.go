package dungeon_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDungeon(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dungeon Suite")
}
