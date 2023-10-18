func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
    root, isBT:=isValidTree(n,leftChild,rightChild)
    if !isBT {
        return false
    }
    vis:=make([]bool,n,n)
    stack:=[]int{root}

    for len(stack)>0 {
        curr:=stack[len(stack)-1]
        stack=stack[:len(stack)-1]

        if vis[curr]{
            return false
        }
        vis[curr]=true
        ln,rn:=leftChild[curr], rightChild[curr]
        if ln!=-1{
            stack=append(stack,ln)
        }
        if rn!=-1{
            stack=append(stack,rn)
        }
    }
    for _,val:=range vis{
        if val==false{
            return false
        }
    }
    // if len(vis)!=n{
    //     return false
    // }
    return true
}

func isValidTree(n int, lc , rc []int) (int, bool){
    temp:= make([]bool,n,n)
    for _, val := range lc{
        if val==-1{
            continue
        }
        temp[val]=true
    }
    for _, val := range rc{
        if val==-1{
            continue
        }
        temp[val]=true
    }
    root :=-1
    pc:=0
    for i, val :=range temp{
        if !val{
            root=i
            pc++
        }
    }
    if pc!=1 || root == -1{
        return -1,false
    }
    return root, true
}
