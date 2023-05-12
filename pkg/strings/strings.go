package strings


func FirstN(s string, n int) string {
	if len(s) > n {
		return s[:n]
	}
	return s
}

// RepeatJoin repeats the string s count times and joins them with seq.
// It returns an empty string if count is negative or zero.
// It panics if the result is longer than the maximum string length.
// It is a combination of strings.Repeat and strings.Join functions.
func RepeatJoin(s string, count int, seq string) string {
	
	// Check if count is negative or zero
	if count <= 0 {
		return ""
	}
	// Check if s is empty
	if s == "" {
		return strings.Repeat(seq, count-1)
	}
	// Create a buffer to write the result
	var buf strings.Builder
	// Write s and seq count-1 times
	for i := 0; i < count-1; i++ {
		buf.WriteString(s)
		buf.WriteString(seq)
	}
	// Write s one more time
	buf.WriteString(s)
	// Return the buffer as a string
	return buf.String()
}
