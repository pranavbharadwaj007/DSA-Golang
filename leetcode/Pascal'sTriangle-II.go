func getRow(row int) []int {
    arr:=make([]int,0)
    val:=1
    for i:=0;i<row;i++{
        arr=append(arr,val)
         val=val*(row-i)/(i+1)
    }
     arr=append(arr,val)
    return arr
}
