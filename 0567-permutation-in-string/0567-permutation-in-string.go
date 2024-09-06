func checkInclusion(s1 string, s2 string) bool {
    leftIndex, count := 0, [26]int{}
    for i := range s1 { 
        count[s1[i]-'a']++ // Using 'a' instead of 97 for clarity
    }

    for rightIndex := range s2 { // Renamed r to rightIndex
        count[s2[rightIndex]-'a']-- // Using 'a' instead of 97 for clarity
        if count == [26]int{} { 
            return true 
        }
        
        if rightIndex + 1 >= len(s1) { // Check if the window size exceeds s1's length
            count[s2[leftIndex]-'a']++ // Using 'a' instead of 97 for clarity
            leftIndex++ // Move the left pointer
        }
    }
    return false
}