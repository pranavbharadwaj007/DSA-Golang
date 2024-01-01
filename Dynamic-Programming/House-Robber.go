
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
func rob(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    } else if n == 1 {
        return nums[0]
    }

    prev1 := nums[0]
    prev2 := 0

    for i := 1; i < n; i++ {
        pick := nums[i]
        if i > 1 {
            pick += prev2
        }
        notPick := 0 + prev1
        curr := max(pick, notPick)
        prev2 = prev1
        prev1 = curr
    }

    return prev1
}

