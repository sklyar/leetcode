package shuffle_the_array

func shuffle(nums []int, n int) []int {
	res := make([]int, len(nums))

	for i := 0; i < n; i++ {
		res[i*2] = nums[i]
		res[i*2+1] = nums[i+n]
	}

	return res
}
