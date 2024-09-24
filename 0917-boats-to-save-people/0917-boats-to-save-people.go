func numRescueBoats(people []int, limit int) int {
    sort.Ints(people)
    res, left, right := 0, 0, len(people)-1
    for left < right {
        if people[left] + people[right] <= limit {
            res++
            left++
        }else {
            res++
        }
        right--
    }
    //if [2,3] and finaly right and left fallen on 2, it means 2 needs own ship, hence res++
    if left == right{
        res++
    }

    return res
}