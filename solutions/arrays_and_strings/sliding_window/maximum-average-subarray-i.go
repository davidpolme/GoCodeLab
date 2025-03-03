package sliding_window

// FindMaxAverage finds the maximum average of a contiguous subarray of length k.
func FindMaxAverage(nums []int, k int) float64 {
    currentSum := 0
    for i := 0; i < k; i++ {
        currentSum += nums[i]
    }

    maxSum := currentSum

    for i := k; i < len(nums); i++ {
        currentSum += nums[i] - nums[i-k]
        if currentSum > maxSum {
            maxSum = currentSum
        }
    }

    return float64(maxSum) / float64(k)
}
