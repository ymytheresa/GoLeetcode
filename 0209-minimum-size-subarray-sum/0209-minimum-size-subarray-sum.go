func minSubArrayLen(target int, nums []int) int {
    res, sum, startIdx := len(nums)+1, 0, 0
    for i, n := range nums{
        sum += n
        for sum >= target{
            res = min(res, i-startIdx+1)
            sum -= nums[startIdx]
            startIdx++
        }
    }

    if res == len(nums)+1 {
        return 0
    }else {
        return res
    }
}