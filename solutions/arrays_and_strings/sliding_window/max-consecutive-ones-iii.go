package sliding_window

// LongestOnes finds the maximum number of consecutive 1s in the array 
// allowing at most k flips from 0 to 1.
func LongestOnes(nums []int, k int) int {
	left, right := 0, 0
	zerosInArray := 0
	maxLength := 0

	for right < len(nums) {
		if nums[right] == 0 {
			zerosInArray++
		}

		for zerosInArray > k {
			if nums[left] == 0 {
				zerosInArray--
			}
			left++
		}

		maxLength = max(maxLength, right-left+1)
		right++
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
