import (
    
    "sort"
)

func minOperations(A []int) int {
    n := len(A)
    ans := n
    hashSet := make(map[int]bool)

    for _, x := range A {
        hashSet[x] = true
    }

    unique := make([]int, 0, len(hashSet))
    for x := range hashSet {
        unique = append(unique, x)
    }

    sort.Ints(unique)

    j := 0
    m := len(unique)

    for i := 0; i < m; i++ {
        for j < m && unique[j] < unique[i]+n {
            j++
        }
        ans = min(ans, n-j+i)
    }

    return ans
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
