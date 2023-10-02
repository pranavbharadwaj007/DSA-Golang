import (
    "strings"
)
func isPalindrome(s string) bool {
    res := strings.ToLower(s)
    tempser := ""
    for _, c := range res {
        if isValid(c) {
        tempser += string(c)
        }
    }
    newChar := reverse(tempser)
    return newChar == tempser
}
func isValid(c rune) bool {
    if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z' ) || (c>= '0' && c<='9') {
        return true
    }
    return false
}
