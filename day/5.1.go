package day

import (
	"advent-of-gode/shared"
	"fmt"
	"strconv"
	"strings"
)

type FreshRange struct {
	start int
	end   int
}

func (fr FreshRange) IsInRange(id int) bool {
	return id >= fr.start && id <= fr.end
}

func NewFreshRangeFromString(input string) *FreshRange {
	split := strings.Split(input, "-")
	s, _ := strconv.Atoi(split[0])
	e, _ := strconv.Atoi(split[1])
	return &FreshRange{start: s, end: e}
}

type IngredientDatabase struct {
	freshRanges []FreshRange
	inventory   []int
}

func NewIngredientDatabaseFromInput(input []string) *IngredientDatabase {
	var fr []FreshRange
	var inv []int

	inventoryMode := false
	for _, line := range input {
		if line == "" {
			inventoryMode = true
		} else if inventoryMode {
			id, _ := strconv.Atoi(string(line))
			inv = append(inv, id)
		} else {
			r := NewFreshRangeFromString(string(line))
			fr = append(fr, *r)
		}
	}

	return &IngredientDatabase{freshRanges: fr, inventory: inv}
}

func (db IngredientDatabase) isIngredientFresh(id int) bool {
	//fmt.Printf("isIngredientFresh(%d)...", id)
	for _, fr := range db.freshRanges {
		if fr.IsInRange(id) {
			//fmt.Printf("true!\n")
			return true
		}
	}

	//fmt.Printf("false!\n")
	return false
}

func (db IngredientDatabase) countFreshInventory() int {
	count := 0

	for _, id := range db.inventory {
		if db.isIngredientFresh(id) {
			count++
		}
	}
	return count
}

func FiveOne() {
	input := shared.GetPuzzleInput("5.txt", "\n")

	//fmt.Println(input)

	/*
		3-5
		10-14
		16-20
		12-18

		1
		5
		8
		11
		17
		32
	*/

	db := NewIngredientDatabaseFromInput(input)

	fmt.Printf("Your inventory contains %d fresh ingredients\n", db.countFreshInventory())
}
