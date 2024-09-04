func trap(height []int) int {
    l, r, lmax, rmax, res := 0, len(height)-1, 0, 0, 0
    
    for l < r {
        lv, rv := height[l], height[r]
        if lv < rv{
            if lv >= lmax {
                lmax = lv
            }else{
                res += lmax - lv
            }
            l++
        }else{
            if rv >= rmax{
                rmax = rv
            }else{
                res += rmax - rv
            }
            r--
        }
    }
    return res
}