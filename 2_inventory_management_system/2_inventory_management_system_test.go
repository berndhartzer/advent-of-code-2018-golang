package two_inventory_management_system

import (
	"testing"
	"os"
	"log"
	"bufio"
)

func TestInventoryManagementSystemPartOneExampleInput(t *testing.T) {
	input := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}
	expected := 12
	output := InventoryManagementSystemPartOne(input)

	if output != expected {
		t.Fail()
	}
}

func TestInventoryManagementSystemPartOneActualInput(t *testing.T) {
	file, err := os.Open("./2_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	output := InventoryManagementSystemPartOne(input)
	t.Log(output)
}

func TestInventoryManagementSystemPartTwoExampleInput(t *testing.T) {
	input := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}
	expected := "fgij"
	output := InventoryManagementSystemPartTwo(input)
	t.Log(output)

	if output != expected {
		t.Fail()
	}
}

func TestInventoryManagementSystemPartTwoActualInput(t *testing.T) {
	file, err := os.Open("./2_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	output := InventoryManagementSystemPartTwo(input)
	t.Log(output)
}
