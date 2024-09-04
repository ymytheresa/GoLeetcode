import "math"

func findMaxAverage(nums []int, k int) float64 {
	max := math.Inf(-1) //infinite negative
	start, sum := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if i-start+1 == k {
			max = math.Max(float64(sum)/float64(k), max)
			sum -= nums[start]
			start++
		}
	}
	return max
}