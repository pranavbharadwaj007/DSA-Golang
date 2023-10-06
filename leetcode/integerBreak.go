func integerBreak(n int) int {
    if n == 2 {
        return 1
    }
    if n == 3 {
        return 2
    }
    threes := n / 3
    rem := n % 3

    if rem == 0 {
        return int(math.Pow(3, float64(threes)))
    } else if rem == 1 {
        return int(math.Pow(3, float64(threes-1))) * 4
    } else {
        return int(math.Pow(3, float64(threes))) * 2
    }
}
