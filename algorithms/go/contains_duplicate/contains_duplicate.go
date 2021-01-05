package contains_duplicate

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{}, len(nums))

	for _, n := range nums {
		_, exist := m[n]
		if exist {
			return true
		}

		m[n] = struct{}{}
	}

	return false
}
