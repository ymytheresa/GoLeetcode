func maximumUniqueSubarray(nums []int) int {
    res, start, sum := 0, 0, 0
    counter := make(map[int]struct{})
    for end, n := range nums{
        if _, found := counter[n] ; found {
            for nums[start] != n && start < end{
                sum -= nums[start]
                delete(counter, nums[start])
                start++
            }
            sum -= n
            start++
        }
        
        sum += n
        res = max(res, sum)
        counter[n] = struct{}{}
    }

    return res
}