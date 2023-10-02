func winnerOfGame(colors string) bool {
    cnta := 0
    cntb := 0
    for i:=1;i< len(colors)-1;i++{
        if colors[i-1] == colors[i] && colors[i]==colors[i+1]{
            if colors[i]=='A'{
                cnta++
            }else{
                cntb++
            }
        }
    }
    return cnta>cntb
}
