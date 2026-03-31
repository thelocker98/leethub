func removeDuplicates(s string) string {
	// Stack
	var stack []rune
	for _, s := range s {
		if len(stack) != 0 && s == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s)
		}
	}
	return string(stack)
}