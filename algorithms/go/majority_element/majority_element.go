package majority_element

// Brute-force
func majorityElement_A(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := make(map[int]int, len(nums)/2)
	for _, n := range nums {
		m[n]++
	}

	maxKey := nums[0]
	for i := 1; i < len(nums); i++ {
		if m[nums[i]] > m[maxKey] {
			maxKey = nums[i]
		}
	}

	return maxKey
}

// Boyer–Moore majority vote algorithm
func majorityElement_B(nums []int) int {
	var (
		candidate int
		count     int
	)

	for _, n := range nums {
		if n == candidate {
			count++
			continue
		}

		if count == 0 {
			candidate = n
			count = 1
		} else {
			count--
		}
	}

	return candidate
}
