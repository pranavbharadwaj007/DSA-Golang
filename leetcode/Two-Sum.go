func twoSum(nums []int, target int) []int {
     mp := map[int]int{}
     var array=make([]int,2)
    for  i,n := range nums{
        tem := target-n
        idx,ok := mp[tem]
        if ok {
            array[0] = idx
            array[1] = i
            return array
        } else {
            mp[n] = i
        }

    }
    return array
}
