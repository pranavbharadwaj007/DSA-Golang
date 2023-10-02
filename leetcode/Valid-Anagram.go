func isAnagram(s string, t string) bool {
    if len(s)!=len(t){
        return false
    }
    cnt := [26]int{}
    for k := range s{
        cnt[s[k]-'a']++
        cnt[t[k]-'a']--
        }  
        for k:=1; k<26; k++{
        if cnt[k]!=0{
            return false;
        }
    }
    return true; 
}
