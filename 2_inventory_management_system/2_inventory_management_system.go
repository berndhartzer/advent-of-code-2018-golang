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
