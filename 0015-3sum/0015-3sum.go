import (
    "sort"
)

func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    res := make([][]int,0)
    for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i-1]{ // prevent looking for k repeatedly
            continue
        }
        // Find all pairs that sum to -nums[i]
        pairs := findAllPairsInSorted(nums[i+1:], -nums[i]) // k should not be inside
        for _, pair := range pairs {
            triplet := []int{nums[i], pair[0], pair[1]} // Construct the triplet
            res = append(res, triplet) // Add the triplet to the result
        }
        
    }
    return res
}

//repeated k/target was taken care, we have to take care if repeated i and j
func findAllPairsInSorted(n []int, k int) [][]int {
    l, r := 0, len(n)-1
    pairs := make([][]int, 0)
    for l < r {
        sum := n[l] + n[r]
        // if sum == k {
        //     return []int{n[l], n[r], -k}, true
        // } 
        // we do not return here coz the same -k might be sum of different pair of i, j

        if sum == k {
            res := []int{n[l], n[r], -k}
            for l < r && n[l] == n[l+1]{
                l++
            }
            for l < r && n[r] == n[r-1]{
                r--
            }
            l++
            r--
            pairs = append(pairs, res)
        }else if sum > k{
            r--
        }else {
            l++
        }
    }
    return pairs
}