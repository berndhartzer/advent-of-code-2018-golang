package two_inventory_management_system

func InventoryManagementSystemPartOne(boxIds []string) int {
	twos := 0
	threes := 0

	for _, id := range boxIds {
		hasTwo := false
		hasThree := false
		counts := make(map[string]int)

		for _, letter := range id {
			counts[string(letter)] += 1
		}

		for _, count := range counts {
			if count == 2 {
				hasTwo = true
			}
			if count == 3 {
				hasThree = true
			}
		}

		if hasTwo {
			twos += 1
		}
		if hasThree {
			threes += 1
		}
	}

	return twos * threes
}

func InventoryManagementSystemPartTwo(boxIds []string) string {
	var chars = make([]byte, len(boxIds[0]) - 1)

	for i, id := range boxIds {

		// We've got our answer, stop looping
		if chars[0] != 0 {
			break
		}

		remaining := boxIds[i + 1:]
		for _, compare := range remaining {
			var common []byte

			for j := 0; j < len(id); j++ {
				if id[j] == compare[j] {
					common = append(common, id[j])
				}

				// This can't be the correct pair
				if len(common) <= j - 1 {
					break
				}
			}

			if len(common) == len(id) - 1 {
				copy(chars, common)
				break
			}
		}
	}

	return string(chars)
}
