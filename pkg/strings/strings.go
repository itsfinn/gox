package strings


func FirstN(s string, n int) string {
	if len(s) > n {
		return s[:n]
	}
	return s
}
