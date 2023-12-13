func numSpecial(mat [][]int) int {
     m,n:=len(mat), len(mat[0])
    row:=make([]int,m)
    col:=make([]int,n)
    for ind,_ := range mat{
        for j,_ := range mat[ind]{
            if mat[ind][j]==1{
                row[ind]++
                col[j]++
            }
        }
    }
    var res=0
    for ind,_ := range mat{
        for j,_ := range mat[ind]{
            if mat[ind][j]==1{
               if row[ind]==1{
                   if col[j]==1{
                       res++
                   }
               }
             
            }
        }
    }
    return res
}
