func removeDuplicates(s string) string {
	// Stack
	var stack []rune
	for _, s := range s {
		if len(stack) != 0 {
			// Pop top
			top := stack[len(stack)-1]
			if s == top {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, s)
			}
		} else {
			stack = append(stack, s)
		}
	}
	return string(stack)
}

