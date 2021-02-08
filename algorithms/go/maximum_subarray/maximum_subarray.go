package maximum_subarray

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	current := nums[0]
	max := current

	for i := 1; i < len(nums); i++ {
		if current < 0 {
			current = nums[i]
		} else {
			current += nums[i]
		}

		if current > max {
			max = current
		}
	}

	return max
}
