func sortByBits(arr []int) []int {
    sort.Slice(arr, func(i, j int) bool {
        x := f(arr[i])
        y := f(arr[j])
        if x != y{
            return x<y
        }
        return arr[i]<arr[j]
    })
    return arr
}

func f(x int) int {
    tmp := x
    ans := 0
    for tmp>0{
        ans+=1
        tmp&=(tmp-1)
    }
    return ans
}
