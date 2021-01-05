package single_number

func singleNumber(nums []int) int {
	var x int

	for i := 0; i < len(nums); i++ {
		x ^= nums[i]
	}

	return x
}
