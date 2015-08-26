package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

var cli string

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var _ = BeforeSuite(func() {
	cli = buildCli()
})

func buildCli() string {
	cliPath, err := gexec.Build("github.com/tscolari/dungeon_game")
	Expect(err).ToNot(HaveOccurred())
	return cliPath
}
