package zebra

import (
	"strings"
)

type Solution struct {
	DrinksWater string
	OwnsZebra   string
}

type Person struct {
	House int
	Color string
	Nationality string
	Drink string
	Smoke string
	Pet string
}

func SolvePuzzle() Solution {
	text := `1. There are five houses.
2. The Englishman lives in the red house.
3. The Spaniard owns the dog.
4. Coffee is drunk in the green house.
5. The Ukrainian drinks tea.
6. The green house is immediately to the right of the ivory house.
7. The Old Gold smoker owns snails.
8. Kools are smoked in the yellow house.
9. Milk is drunk in the middle house.
10. The Norwegian lives in the first house.
11. The man who smokes Chesterfields lives in the house next to the man with the fox.
12. Kools are smoked in the house next to the house where the horse is kept.
13. The Lucky Strike smoker drinks orange juice.
14. The Japanese smokes Parliaments.
15. The Norwegian lives next to the blue house.`
	listOfText := strings.Split(text,"\n")
	colors := getColors()
	nationality := getNationality()
	drink := getDrink()
	smoke := getSmoke()
	pets := getPet()

	var table []Person
	for _, t := range listOfText {

		p := Person{}

		//	check colors
		for _, c := range colors {
			if contain(t, c) {
				p.Color = c
			}
		}

		// check nationality
		for _, n := range nationality {
			if contain(t, n) {
				p.Nationality = n
			}
		}


		// check drink
		for _, d := range drink {
			if contain(t, d) {
				p.Drink = d
			}
		}

		// check smoke
		for _, s := range smoke {
			if contain(t, s) {
				p.Smoke = s
			}
		}

		// check pet
		for _, pet := range pets {
			if contain(t, pet) {
				p.Pet = pet
			}
		}

		table = append(table, p)

	}
	return Solution{}
}

func getColors() []string {
	return []string{"Yellow", "Blue", "Red", "Ivory", "Green"}
}

func getNationality() []string {
	return []string{"Norwegian", "Ukrainian", "Englishman", "Spaniard", "Japanese"}
}

func getDrink() []string {
	return []string{"Water", "Tea", "Milk", "Orange juice", "Coffee"}
}

func getSmoke() []string {
	return []string{"Kools", "Chesterfield", "Old Gold", "Lucky Strike", "Parliament"}
}

func getPet() []string {
	return []string{"Fox", "Horse", "Snails", "Dog", "Zebra"}
}

func contain (text string, s string) bool {
	newTexts := strings.Split(text," ")
	for _,v := range newTexts {
		if v == s {
			return true
		}
	}
	return false
}
