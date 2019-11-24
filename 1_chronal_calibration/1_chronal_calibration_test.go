package one_chronal_calibration

import (
	"testing"
	"os"
	"log"
	"bufio"
)

func TestChronalCalibrationPartOneActualInput(t *testing.T) {
	file, err := os.Open("./1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	output := ChronalCalibrationPartOne(input)
	t.Log(output)
}

func TestChronalCalibrationPartTwoActualInput(t *testing.T) {
	file, err := os.Open("./1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	output := ChronalCalibrationPartTwo(input)
	t.Log(output)
}
