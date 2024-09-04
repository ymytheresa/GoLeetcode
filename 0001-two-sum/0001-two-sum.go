func twoSum(nums []int, target int) []int {
    counter := map[int]int{}
	for i := 0; i < len(nums); i++ {
        another := target - nums[i]
        if _, found := counter[another]; found{
            return []int{counter[another], i}
        }
        counter[nums[i]] = i
	}
    return nil
}