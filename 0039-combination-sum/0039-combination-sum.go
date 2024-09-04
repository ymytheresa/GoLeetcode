func combinationSum(candidates []int, target int) [][]int {
   sort.Ints(candidates)
   combinations := []int{}
   result := [][]int{}
   backtrack(candidates, combinations, target, 0, &result)
   return result
}

func backtrack(cand []int, comb []int, t int, idx int, res *[][]int){
    if t == 0{
        tempComb := make([]int, len(comb))
        copy(tempComb, comb)
        *res = append(*res, tempComb)
        return
    }else if t < 0{
        return
    }else{
        for idx < len(cand){
            backtrack(cand, append(comb, cand[idx]), t-cand[idx], idx, res)
            idx++
        }
    }
}