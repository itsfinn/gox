package strings

import "testing"

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


func TestJoinNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		sep      string
		expected string
	}{
		// int tests
		{
			name:     "empty int slice",
			input:    []int{},
			sep:      ",",
			expected: "",
		},
		{
			name:     "single int element",
			input:    []int{42},
			sep:      ",",
			expected: "42",
		},
		{
			name:     "multiple int elements",
			input:    []int{1, 2, 3},
			sep:      ",",
			expected: "1,2,3",
		},
		{
			name:     "int with different separator",
			input:    []int{1, 2, 3},
			sep:      " - ",
			expected: "1 - 2 - 3",
		},
		// int8 tests
		{
			name:     "empty int8 slice",
			input:    []int8{},
			sep:      ",",
			expected: "",
		},
		{
			name:     "single int8 element",
			input:    []int8{42},
			sep:      ",",
			expected: "42",
		},
		{
			name:     "multiple int8 elements",
			input:    []int8{1, 2, 3},
			sep:      ",",
			expected: "1,2,3",
		},

		// int16 tests
		{
			name:     "empty int16 slice",
			input:    []int16{},
			sep:      ",",
			expected: "",
		},
		{
			name:     "single int16 element",
			input:    []int16{42},
			sep:      ",",
			expected: "42",
		},
		{
			name:     "multiple int16 elements",
			input:    []int16{1, 2, 3},
			sep:      ",",
			expected: "1,2,3",
		},

		// int32 tests
		{
			name:     "empty int32 slice",
			input:    []int32{},
			sep:      ",",
			expected: "",
		},
		{
			name:     "single int32 element",
			input:    []int32{42},
			sep:      ",",
			expected: "42",
		},
		{
			name:     "multiple int32 elements",
			input:    []int32{1, 2, 3},
			sep:      ",",
			expected: "1,2,3",
		},

		// int64 tests
		{
			name:     "empty int64 slice",
			input:    []int64{},
			sep:      ",",
			expected: "",
		},
		{
			name:     "single int64 element",
			input:    []int64{42},
			sep:      ",",
			expected: "42",
		},
		{
			name:     "multiple int64 elements",
			input:    []int64{1, 2, 3},
			sep:      ",",
			expected: "1,2,3",
		},

		// float32 tests
		{
			name:     "empty float32 slice",
			input:    []float32{},
			sep:      ",",
			expected: "",
		},
		{
			name:     "single float32 element",
			input:    []float32{3.14},
			sep:      ",",
			expected: "3.14",
		},
		{
			name:     "multiple float32 elements",
			input:    []float32{1.1, 2.2, 3.3},
			sep:      ",",
			expected: "1.1,2.2,3.3",
		},

		// float64 tests
		{
			name:     "empty float64 slice",
			input:    []float64{},
			sep:      ",",
			expected: "",
		},
		{
			name:     "single float64 element",
			input:    []float64{3.1415926},
			sep:      ",",
			expected: "3.1415926",
		},
		{
			name:     "multiple float64 elements",
			input:    []float64{1.1, 2.2, 3.3},
			sep:      ",",
			expected: "1.1,2.2,3.3",
		},

		// uint tests
		{
			name:     "empty uint slice",
			input:    []uint{},
			sep:      ",",
			expected: "",
		},
		{
			name:     "single uint element",
			input:    []uint{42},
			sep:      ",",
			expected: "42",
		},
		{
			name:     "multiple uint elements",
			input:    []uint{1, 2, 3},
			sep:      ",",
			expected: "1,2,3",
		},

		// uint8 tests
		{
			name:     "empty uint8 slice",
			input:    []uint8{},
			sep:      ",",
			expected: "",
		},
		{
			name:     "single uint8 element",
			input:    []uint8{42},
			sep:      ",",
			expected: "42",
		},
		{
			name:     "multiple uint8 elements",
			input:    []uint8{1, 2, 3},
			sep:      ",",
			expected: "1,2,3",
		},

		// uint16 tests
		{
			name:     "empty uint16 slice",
			input:    []uint16{},
			sep:      ",",
			expected: "",
		},
		{
			name:     "single uint16 element",
			input:    []uint16{42},
			sep:      ",",
			expected: "42",
		},
		{
			name:     "multiple uint16 elements",
			input:    []uint16{1, 2, 3},
			sep:      ",",
			expected: "1,2,3",
		},

		// uint32 tests
		{
			name:     "empty uint32 slice",
			input:    []uint32{},
			sep:      ",",
			expected: "",
		},
		{
			name:     "single uint32 element",
			input:    []uint32{42},
			sep:      ",",
			expected: "42",
		},
		{
			name:     "multiple uint32 elements",
			input:    []uint32{1, 2, 3},
			sep:      ",",
			expected: "1,2,3",
		},

		// uint64 tests
		{
			name:     "empty uint64 slice",
			input:    []uint64{},
			sep:      ",",
			expected: "",
		},
		{
			name:     "single uint64 element",
			input:    []uint64{42},
			sep:      ",",
			expected: "42",
		},
		{
			name:     "multiple uint64 elements",
			input:    []uint64{1, 2, 3},
			sep:      ",",
			expected: "1,2,3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result string
			switch v := tt.input.(type) {
			case []int:
				result = JoinNumbers(v, tt.sep)
			case []int8:
				result = JoinNumbers(v, tt.sep)
			case []int16:
				result = JoinNumbers(v, tt.sep)
			case []int32:
				result = JoinNumbers(v, tt.sep)
			case []int64:
				result = JoinNumbers(v, tt.sep)
			case []float32:
				result = JoinNumbers(v, tt.sep)
			case []float64:
				result = JoinNumbers(v, tt.sep)
			case []uint:
				result = JoinNumbers(v, tt.sep)
			case []uint8:
				result = JoinNumbers(v, tt.sep)
			case []uint16:
				result = JoinNumbers(v, tt.sep)
			case []uint32:
				result = JoinNumbers(v, tt.sep)
			case []uint64:
				result = JoinNumbers(v, tt.sep)
			default:
				t.Fatalf("unsupported type: %T", v)
			}

			if result != tt.expected {
				t.Errorf("JoinNumbers() = %v, want %v", result, tt.expected)
			}
		})
	}
}
func TestSplitToNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		sep      string
		typ      interface{}
		expected interface{}
		wantErr  bool
	}{
		// int tests
		{
			name:     "empty string to int",
			input:    "",
			sep:      ",",
			typ:      []int{},
			expected: []int{},
			wantErr:  false,
		},
		{
			name:     "single int",
			input:    "42",
			sep:      ",",
			typ:      []int{},
			expected: []int{42},
			wantErr:  false,
		},
		{
			name:     "multiple ints",
			input:    "1,2,3",
			sep:      ",",
			typ:      []int{},
			expected: []int{1, 2, 3},
			wantErr:  false,
		},
		{
			name:     "invalid int",
			input:    "1,a,3",
			sep:      ",",
			typ:      []int{},
			expected: nil,
			wantErr:  true,
		},

		// int32 tests
		{
			name:     "single int32",
			input:    "42",
			sep:      ",",
			typ:      []int32{},
			expected: []int32{42},
			wantErr:  false,
		},
		{
			name:     "multiple int32s",
			input:    "1,2,3",
			sep:      ",",
			typ:      []int32{},
			expected: []int32{1, 2, 3},
			wantErr:  false,
		},

		// uint32 tests
		{
			name:     "single uint32",
			input:    "42",
			sep:      ",",
			typ:      []uint32{},
			expected: []uint32{42},
			wantErr:  false,
		},
		{
			name:     "multiple uint32s",
			input:    "1,2,3",
			sep:      ",",
			typ:      []uint32{},
			expected: []uint32{1, 2, 3},
			wantErr:  false,
		},
		{
			name:     "negative uint32",
			input:    "-1",
			sep:      ",",
			typ:      []uint32{},
			expected: nil,
			wantErr:  true,
		},

		// float32 tests
		{
			name:     "single float32",
			input:    "3.14",
			sep:      ",",
			typ:      []float32{},
			expected: []float32{3.14},
			wantErr:  false,
		},
		{
			name:     "multiple float32s",
			input:    "1.1,2.2,3.3",
			sep:      ",",
			typ:      []float32{},
			expected: []float32{1.1, 2.2, 3.3},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result interface{}
			var err error

			switch tt.typ.(type) {
			case []int:
				result, err = SplitToNumbers[int](tt.input, tt.sep)
			case []int32:
				result, err = SplitToNumbers[int32](tt.input, tt.sep)
			case []uint32:
				result, err = SplitToNumbers[uint32](tt.input, tt.sep)
			case []float32:
				result, err = SplitToNumbers[float32](tt.input, tt.sep)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("SplitToNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !equalSlices(result, tt.expected) {
				t.Errorf("SplitToNumbers() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Helper function to compare slices of different types
func equalSlices(a, b interface{}) bool {
	switch va := a.(type) {
	case []int:
		vb := b.([]int)
		if len(va) != len(vb) {
			return false
		}
		for i := range va {
			if va[i] != vb[i] {
				return false
			}
		}
		return true
	case []int32:
		vb := b.([]int32)
		if len(va) != len(vb) {
			return false
		}
		for i := range va {
			if va[i] != vb[i] {
				return false
			}
		}
		return true
	case []uint32:
		vb := b.([]uint32)
		if len(va) != len(vb) {
			return false
		}
		for i := range va {
			if va[i] != vb[i] {
				return false
			}
		}
		return true
	case []float32:
		vb := b.([]float32)
		if len(va) != len(vb) {
			return false
		}
		for i := range va {
			if va[i] != vb[i] {
				return false
			}
		}
		return true
	default:
		return false
	}
}

