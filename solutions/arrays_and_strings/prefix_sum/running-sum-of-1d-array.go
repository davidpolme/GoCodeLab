package prefix_sum

func RunningSum(nums []int) []int {
    prefix := make([]int, len(nums))

    for i := 0; i < len(nums); i++ {
        if i == 0 {
            prefix[i] = nums[i]
        } else {
            prefix[i] = prefix[i-1] + nums[i]
        }
    }
    return prefix
}
