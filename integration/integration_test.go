package integration_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Integration", func() {

	Context("when things go wrong", func() {
		It("complains if no dungeon file is given", func() {
			cmd := exec.Command(cli)

			output, err := cmd.CombinedOutput()
			Expect(err).To(HaveOccurred())
			Expect(output).To(MatchRegexp("Usage of"))
		})

		It("complains if there's issues loading the dungeon file", func() {
			cmd := exec.Command(cli, "-dungeon", "../dungeon/fixtures/invalid_dungeon-not_enough_rooms.yml")

			output, err := cmd.CombinedOutput()
			Expect(err).To(HaveOccurred())
			Expect(output).To(MatchRegexp("Invalid dungeon file"))
		})
	})

	It("doesn't fail to run", func() {
		cmd := exec.Command(cli, "-dungeon", "../dungeon/fixtures/valid_dungeon.yml")

		err := cmd.Run()
		Expect(err).ToNot(HaveOccurred())
	})

})
