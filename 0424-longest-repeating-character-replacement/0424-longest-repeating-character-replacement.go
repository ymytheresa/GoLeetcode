func characterReplacement(s string, k int) int {
    res, left, maxf := 0, 0, 0
    counter := [26]int{}

    for end, r := range s {
        counter[r-'A']++
        maxf = max(maxf, counter[r-'A'])
        if (end-left+1) - maxf > k {
            counter[s[left]-'A']--
            left ++
        }
        res = max(res, end-left+1)
    }



    return res
}