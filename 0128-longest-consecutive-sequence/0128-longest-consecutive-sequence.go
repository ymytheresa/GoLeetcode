// longestConsecutive finds the length of the longest consecutive elements sequence.
func longestConsecutive(nums []int) int {
    set := make(map[int]struct{})
    for _, n := range nums{
        set[n] = struct{}{}
    }

    max := 0
    for n, _ := range set{
        //if we found n-1 we can continue
        //coz the later part will ensure that we (n) must have been seen before
        if _, found := set[n-1]; !found{
            // if n-1 doesnt exist, it means n is the new boundary, then we start looking for n+1...
            nPlus := n+1
            currentStreak := 1
            for _, found := set[nPlus]; found; _, found = set[nPlus]{
                nPlus++
                currentStreak++
            }
            max = maxInt(max, currentStreak)
        }
    }

    return max
}

func maxInt(a, b int) int{
    if a > b{
        return a
    }
    return b
}