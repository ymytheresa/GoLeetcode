func lengthOfLongestSubstring(s string) int {
    counter := make(map[rune]int) //met alphabet, idx
    res, start := 0, 0
    for end, r := range s {
        if _, found := counter[r]; !found{
            res = max(res, end-start+1)
            counter[r] = end
            continue
        }
        //dup
        if counter[r]+1 >= start{
            start = counter[r]+1
        }
        // start = counter[r]+1
        res = max(res, end-start+1)
        counter[r] = end
    }
    return res
}


// func lengthOfLongestSubstring(s string) int {
//     charIndex := make(map[rune]int)
//     maxLength, start := 0, 0

//     for end, char := range s {
//         if lastIndex, found := charIndex[char]; found && lastIndex >= start {
//             start = lastIndex + 1
//         }
//         charIndex[char] = end
//         maxLength = max(maxLength, end - start + 1)
//     }

//     return maxLength
// }