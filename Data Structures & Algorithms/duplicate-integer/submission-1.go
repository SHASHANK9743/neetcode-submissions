func hasDuplicate(nums []int) bool {
    freq_map := make(map[int]int)
    for _, v := range nums{
        _, ok := freq_map[v]
        if ok {
            return true
        } else {
            freq_map[v] = 1
        }
    }
    return false
}
