package two_pointers

func abs(num int) int {
    if num < 0 {
        return -num
    }
    return num
}

func square(num int) int {
    return num * num
}

// SortedSquares returns a sorted array of the squares of a sorted array.
func SortedSquares(nums []int) []int {
    numsLen := len(nums)
    left, right := 0, numsLen-1
    result := make([]int, numsLen)

    for index := numsLen - 1; left <= right; index-- {
        if abs(nums[left]) > abs(nums[right]) {
            result[index] = square(nums[left])
            left++
        } else {
            result[index] = square(nums[right])
            right--
        }
    }
    return result
}
