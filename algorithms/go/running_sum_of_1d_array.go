package _go

func runningSum(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	res := make([]int, len(nums))
	res[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] + nums[i]
	}

	return res
}
