func hasDuplicate(nums []int) bool {
    for i, iValue := range nums {
        for j := i + 1; j < len(nums); j++ {
            if iValue == nums[j] {
                return true
            }
        }
    }

    return false
}
