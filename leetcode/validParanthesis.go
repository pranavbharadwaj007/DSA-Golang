func isValid(s string) bool {
     stack := make([]byte, 0)
    closeToOpen := map[byte]byte{')': '(', ']': '[', '}': '{'}
    for _, c := range []byte(s) {
        open, ok := closeToOpen[c]
        if ok {
            if len(stack) == 0 || stack[len(stack) - 1] != open {
                return false
            } else {
                stack = stack[:len(stack) - 1]
            }
        } else {
            stack = append(stack, c)
        }
    }
    return len(stack) == 0
}
