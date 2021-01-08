package top_k_frequent_elements

// bucket sort
func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, n := range nums {
		counts[n]++
	}

	buckets := make([][]int, len(nums)+1)
	for frequent, count := range counts {
		buckets[count] = append(buckets[count], frequent)
	}

	res := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0; i-- {
		for j := 0; j < len(buckets[i]); j++ {
			res = append(res, buckets[i][j])
			if len(res) == k {
				return res
			}
		}
	}

	return res
}
