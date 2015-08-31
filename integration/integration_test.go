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
		cmd := exec.Command(cli, "-dungeon", "./fixtures/dungeon.yml")

		err := cmd.Run()
		Expect(err).ToNot(HaveOccurred())
	})

	It("gives the correct directions", func() {
		cmd := exec.Command(cli, "-dungeon", "./fixtures/dungeon.yml")

		output, err := cmd.Output()
		Expect(err).ToNot(HaveOccurred())

		Expect(string(output)).To(MatchRegexp("BEST ROUTE: RIGHT -> DOWN -> DOWN -> RIGHT -> RIGHT"))
	})

	It("gives the correct life points needed", func() {
		cmd := exec.Command(cli, "-dungeon", "./fixtures/dungeon.yml")

		output, err := cmd.Output()
		Expect(err).ToNot(HaveOccurred())

		Expect(string(output)).To(MatchRegexp("MIN HP: 4"))
	})

})
