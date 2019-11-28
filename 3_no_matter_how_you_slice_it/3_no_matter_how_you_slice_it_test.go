package three_no_matter_how_you_slice_it

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestNoMatterHowYouSliceItPartOneExampleInput(t *testing.T) {
	input := []string{
		"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2",
	}
	expected := 4
	output := NoMatterHowYouSliceItPartOne(input)

	if output != expected {
		t.Fail()
	}
}

func TestNoMatterHowYouSliceItPartOneActualInput(t *testing.T) {
	file, err := os.Open("./3_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	output := NoMatterHowYouSliceItPartOne(input)
	t.Log(output)
}

func TestNoMatterHowYouSliceItPartTwoExampleInput(t *testing.T) {
	input := []string{
		"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2",
	}
	expected := 3
	output := NoMatterHowYouSliceItPartTwo(input)

	if output != expected {
		t.Fail()
	}
}

func TestNoMatterHowYouSliceItPartTwoActualInput(t *testing.T) {
	file, err := os.Open("./3_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	output := NoMatterHowYouSliceItPartTwo(input)
	t.Log(output)
}
