package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
)

var numOfGoals = 1
var maxYumLevel = 50
var includeLockedFoods = false
var enableStressMode = false
var displayYumLevel = false

func init() {
	flag.IntVar(&numOfGoals, "num-goals", numOfGoals, "Number of goals to generate")
	flag.IntVar(&numOfGoals, "n", numOfGoals, "Number of goals to generate")
	flag.IntVar(&maxYumLevel, "yum", maxYumLevel, "Max yum level")
	flag.IntVar(&maxYumLevel, "y", maxYumLevel, "Max yum level")
	flag.BoolVar(&includeLockedFoods, "include-locked", includeLockedFoods, "Include locked foods")
	flag.BoolVar(&includeLockedFoods, "l", includeLockedFoods, "Include locked foods")
	flag.BoolVar(&enableStressMode, "stress", enableStressMode, "Enable stress mode")
	flag.BoolVar(&enableStressMode, "s", enableStressMode, "Enable stress mode")
	flag.BoolVar(&displayYumLevel, "display-yum", displayYumLevel, "Display yum level requirements")
	flag.BoolVar(&displayYumLevel, "d", displayYumLevel, "Display yum level requirements")
}

func main() {
	flag.Parse()
	for range numOfGoals {
		GenerateGoal(maxYumLevel, includeLockedFoods, enableStressMode, displayYumLevel)
	}
}

func GenerateGoal(maxYumLevel int, includeLockedFoods, enableStressMode, displayYumLevel bool) {
	goalType := rand.Intn(2)
	switch goalType {
	case 0:
		fmt.Println(chefForHireGoal(maxYumLevel, displayYumLevel))
	case 1:
		fmt.Println(csdGoal(maxYumLevel, includeLockedFoods, enableStressMode))
	}
}

type chefForHireShift struct {
	restaurant string
	shift      int
	yumLevel   int
}

func getEligibleShifts(maxYumLevel int) []chefForHireShift {
	var shifts []chefForHireShift
	for restaurant, levels := range chefForHireLevels {
		for shiftNumber, level := range levels {
			if level <= maxYumLevel {
				shifts = append(shifts, chefForHireShift{restaurant, shiftNumber, level})
			}
		}
	}
	return shifts
}

func chefForHireGoal(maxYumLevel int, displayYumLevel bool) string {
	shifts := getEligibleShifts(maxYumLevel)
	shift := shifts[rand.Intn(len(shifts))]

	goal := fmt.Sprintf("Get a Perfect Day in %s - Shift %d", shift.restaurant, shift.shift+1)
	if displayYumLevel {
		goal += fmt.Sprintf(" (Yum Level %d)", shift.yumLevel)
	}
	return goal
}

func csdGoal(maxYumLevel int, includeLockedFoods, enableStressMode bool) string {
	var maxEntrees, maxSides, maxDrinks int

	switch {
	case maxYumLevel < 4:
		maxEntrees = 3
	case maxYumLevel < 14:
		maxEntrees = 4
	case maxYumLevel < 18:
		maxEntrees = 5
	default:
		maxEntrees = 6
	}

	switch {
	case maxYumLevel < 7:
		maxSides = 1
	case maxYumLevel < 21:
		maxSides = 2
	default:
		maxSides = 3
	}

	switch {
	case maxYumLevel < 11:
		maxDrinks = 1
	default:
		maxDrinks = 2
	}

	numEntrees := rand.Intn(maxEntrees) + 1
	numSides := rand.Intn(maxSides) + 1
	numDrinks := rand.Intn(maxDrinks) + 1

	var entrees, sides, drinks []string
	if includeLockedFoods {
		entrees, sides, drinks = allEntrees, allSides, allDrinks
	} else {
		entrees, sides, drinks = unlockedEntrees, unlockedSides, unlockedDrinks
	}

	entreeIndexes := rand.Perm(len(entrees))[:numEntrees]
	sideIndexes := rand.Perm(len(sides))[:numSides]
	drinkIndexes := rand.Perm(len(drinks))[:numDrinks]

	maxMode := len(cookServeDeliciousModes) - 1
	if enableStressMode {
		maxMode++
	}
	modeIndex := rand.Intn(maxMode)
	mode := cookServeDeliciousModes[modeIndex]
	entreeChoices := make([]string, numEntrees)
	for i, index := range entreeIndexes {
		entreeChoices[i] = entrees[index]
	}
	sidesChoices := make([]string, numSides)
	for i, index := range sideIndexes {
		sidesChoices[i] = sides[index]
	}
	drinkChoices := make([]string, numDrinks)
	for i, index := range drinkIndexes {
		drinkChoices[i] = drinks[index]
	}

	return fmt.Sprintf("Get a Perfect Day in Cook Serve Delicious in %s with entrees: [%s], sides: [%s], drinks: [%s]",
		strings.ToUpper(mode), strings.Join(entreeChoices, ", "), strings.Join(sidesChoices, ", "),
		strings.Join(drinkChoices, ", "))
}
