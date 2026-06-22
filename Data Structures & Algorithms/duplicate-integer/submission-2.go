func hasDuplicate(nums []int) bool {

    mySet := make(map[int]struct{})

    for _, iValue := range nums {
        _, exists := mySet[iValue]
        if exists {
            return true
        } else {
            mySet[iValue] = struct{}{}
        }
    }

    return false
}
