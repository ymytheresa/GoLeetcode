func minSubArrayLen(target int, nums []int) int {
    res, sum, startIdx := math.MaxInt32, 0, 0
    found := false
    for i, n := range nums{
        sum += n
        for sum >= target{
            res = min(res, i-startIdx+1)
            found = true
            sum -= nums[startIdx]
            startIdx++
        }
    }

    if found {
        return res
    }else {
        return 0
    }
}