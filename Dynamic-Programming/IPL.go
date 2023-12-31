// https://www.codechef.com/practice/course/dynamic-programming/INTDP01/problems/ZCO14004


package main
import "fmt"

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	}
	return c
}

func supw(n int, arr []int)int{
    dp := make([]int, n)
    dp[0]=arr[0]
    dp[1]=arr[1]
    dp[2]=arr[2]
    
    for i:=3 ;i<n;i++{
        dp[i]=min(dp[i-1],dp[i-2],dp[i-3])+arr[i]
    }
    return min(dp[n-1], dp[n-2], dp[n-3])
}

func main(){
	 var n int
    fmt.Scan(&n)
    arr := make([]int, n)
    var sum int
    sum=0
    for i := 0; i < n; i++ {
        
        fmt.Scan(&arr[i])
        sum+=arr[i]
    }
    results:=supw(n,arr)
    fmt.Println(sum-results)
}
