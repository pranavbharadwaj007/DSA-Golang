func backspaceCompare(s string, t string) bool {
    stk1:=[]rune{}
    stk2:=[]rune{}
    for _,val := range s{
       if val=='#'{
           if len(stk1)>0{
               stk1=stk1[:len(stk1)-1]
           }
       }else{
           stk1=append(stk1,val)
       }
    }
    for _,val := range t{
        if val=='#'{
           if len(stk2)>0{
               stk2=stk2[:len(stk2)-1]
           }
       }else{
           stk2=append(stk2,val)
       }
    }
    if len(stk1)!=len(stk2){
        return false
    }
    for idx:=range stk1{
        if stk1[idx]!=stk2[idx]{
            return false
        }

    }
            return true
}
