package a

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return getBST(nums, 0, len(nums)-1)
}

func getBST(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  getBST(nums, left, mid-1),
		Right: getBST(nums, mid+1, right),
	}
}
