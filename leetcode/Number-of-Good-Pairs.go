func numIdenticalPairs(nums []int) int {
    cnt :=0
    ht :=make(map[int]int)
    for _,v := range nums{
        oc, exists := ht[v]
        if exists{
            cnt+=oc
            ht[v]=oc+1
        }else{
            ht[v]=1
        }
    }
    return cnt
}
