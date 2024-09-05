func totalFruit(trees []int) int {
    treeTypeMap := make(map[int]bool)
    var res int
    var start int

    for end, treeType := range trees {
        if len(treeTypeMap) < 2 && !treeTypeMap[treeType]{
            treeTypeMap[treeType] = true
        } else if treeTypeMap[treeType]{
        } else {
            treeTypeMap = make(map[int]bool)
            treeTypeMap[treeType] = true
            treeTypeMap[trees[end-1]] = true
            start = end-1
            for trees[start] == trees[start-1]{
                start--
            }
        }
        res = max(res, end-start+1)
    }
    return res
}