package strings

// TestRepeatJoin tests the RepeatJoin function with some examples.
func TestRepeatJoin(t *testing.T) {
	// Define some test cases
	cases := []struct {
		s     string
		count int
		seq   string
		want  string
	}{
		{"?", 3, ",", "?,?,?"},
		{"", 5, "-", "----"},
		{"a", 0, "+", ""},
		{"b", 1, "*", "b"},
		{"c", -2, "/", ""},
	}
	// Loop over the test cases
	for _, c := range cases {
		// Call the RepeatJoin function
		got := RepeatJoin(c.s, c.count, c.seq)
		// Check if the result matches the expected value
		if got != c.want {
			t.Errorf("RepeatJoin(%q, %d, %q) = %q; want %q", c.s, c.count, c.seq, got, c.want)
		}
	}
}

// BenchmarkRepeatJoin benchmarks the RepeatJoin function with some examples.
func BenchmarkRepeatJoin(b *testing.B) {
	// Define some test cases
	cases := []struct {
		s     string
		count int
		seq   string
	}{
		{"?", 3, ","},
		{"", 5, "-"},
		{"a", 0, "+"},
		{"b", 1, "*"},
		{"c", -2, "/"},
	}
	// Loop over the test cases
	for _, c := range cases {
		// Run the benchmark for each case
		b.Run(fmt.Sprintf("%s_%d_%s", c.s, c.count, c.seq), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				repeatJoin(c.s, c.count, c.seq)
			}
		})
	}
}
