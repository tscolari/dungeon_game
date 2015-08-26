package dungeon_test

import (
	"github.com/tscolari/dungeon_game/dungeon"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Load", func() {
	Context("when there's an error opening the file", func() {
		It("returns an error", func() {
			_, err := dungeon.Load("/tmp/this_file_doesnt_exist")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(MatchRegexp("Failed to open dungeon file"))
		})
	})

	Context("when the file is not in the correct format", func() {

		Context("when the room rows have different dimensions", func() {
			dungeonFileName := "./fixtures/invalid_dungeon-wrong_dimensions.yml"

			It("returns an invalid file error", func() {
				_, err := dungeon.Load(dungeonFileName)
				Expect(err).To(MatchError("Invalid dungeon file: All rows must have the same size"))
			})
		})

		Context("when the file is not a valid yaml", func() {
			dungeonFileName := "./fixtures/invalid_dungeon-invalid_YAML.yml"
			It("returns an invalid yaml error", func() {
				_, err := dungeon.Load(dungeonFileName)
				Expect(err).To(MatchError("Invalid dungeon file: Not a valid yaml"))
			})
		})

		Context("when there are less than 2 rooms", func() {
			It("returns an not enough rooms error", func() {
				dungeonFileName := "./fixtures/invalid_dungeon-not_enough_rooms.yml"
				_, err := dungeon.Load(dungeonFileName)
				Expect(err).To(MatchError("Invalid dungeon file: Not enough rooms (min 2)"))
			})
		})

	})

	It("loads the file into a correct Dungeon object", func() {
		dungeonFileName := "./fixtures/valid_dungeon.yml"
		dungeon, err := dungeon.Load(dungeonFileName)
		Expect(err).ToNot(HaveOccurred())

		Expect(dungeon.Rooms[0]).To(Equal([]int{1, 2, 3, 4}))
		Expect(dungeon.Rooms[1]).To(Equal([]int{5, 6, 7, 8}))
		Expect(dungeon.Rooms[2]).To(Equal([]int{9, 10, 11, 12}))
	})
})
