
func isValidSudoku(board [][]byte) bool {
    setv := map[string]bool{}
    for i:=0; i<9;i++{
        for j:=0;j<9;j++{
            if board[i][j]!='.'{
                sr := "row" + strconv.Itoa(i) + string(board[i][j])
                sc := "col" + strconv.Itoa(j) + string(board[i][j])
                boxNumber := (i/3)*3 + j/3
                sb := "box" + strconv.Itoa(boxNumber) + string(board[i][j])
            if setv[sr] || setv[sc]{
                return false
            }
            if setv[sb]{
                return false
            }
            setv[sr]=true
            setv[sc]=true
            setv[sb]=true
            }

        }
    }
    return true
}
