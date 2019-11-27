package three_no_matter_how_you_slice_it

import (
	"fmt"
	"regexp"
	"strconv"
)

var fabric map[string]int

type Claim struct {
	ID, X, Y, W, H int
}

func NoMatterHowYouSliceItPartOne(claims []string) int {
	fabric = make(map[string]int)
	getNumsRegex := regexp.MustCompile("[0-9]+")

	for _, raw := range claims {
		var nums = []int{}

		rawNums := getNumsRegex.FindAllString(raw, -1)
		for _, rawNum := range rawNums {
			num, _ := strconv.Atoi(rawNum)
			nums = append(nums, num)
		}

		claim := Claim{
			nums[0],
			nums[1],
			nums[2],
			nums[3],
			nums[4],
		}

		for i := claim.X; i < claim.X + claim.W; i++ {
			for j := claim.Y; j < claim.Y + claim.H; j++ {
				fabric[fmt.Sprintf("%d,%d", i, j)] += 1
			}
		}
	}

	overlaps := 0
	for _, value := range fabric {
		if value > 1 {
			overlaps += 1
		}
	}

	return overlaps
}

func NoMatterHowYouSliceItPartTwo(claims []string) int {
	return 0
}
