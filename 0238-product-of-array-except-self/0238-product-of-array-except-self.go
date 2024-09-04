func productExceptSelf(nums []int) []int {
    res := make([]int, len(nums))

    //left
    left := 1    
	for i := 0; i < len(nums); i++ {
        res[i] = left
        left *= nums[i]
	}

    //right
    right := 1
	for i := len(nums)-1 ; i >= 0; i-- {
        res[i] *= right
        right *= nums[i]
	}

    return res
}