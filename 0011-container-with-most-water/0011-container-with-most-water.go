func maxArea(height []int) int {
    max := 0
    l, r := 0, len(height)-1

    for l < r {
        // Calculate area directly without a separate function
        minHeight := height[l]
        if height[r] < minHeight {
            minHeight = height[r]
        }
        width := r - l
        area := minHeight * width

        if area > max {
            max = area
        }

        // Optimize pointer movement
        for l < r && height[l] <= minHeight {
            l++
        }
        for l < r && height[r] <= minHeight {
            r--
        }
    } 
    return max
}