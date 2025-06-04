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

// JoinNumbers is a generic function to join any numeric type slice
func JoinNumbers[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](elems []T, sep string) string {
	if len(elems) == 0 {
		return ""
	}

	var s string
	switch v := any(elems[0]).(type) {
	case int:
		s = strconv.Itoa(v)
	case int8:
		s = strconv.FormatInt(int64(v), 10)
	case int16:
		s = strconv.FormatInt(int64(v), 10)
	case int32:
		s = strconv.FormatInt(int64(v), 10)
	case int64:
		s = strconv.FormatInt(v, 10)
	case uint:
		s = strconv.FormatUint(uint64(v), 10)
	case uint8:
		s = strconv.FormatUint(uint64(v), 10)
	case uint16:
		s = strconv.FormatUint(uint64(v), 10)
	case uint32:
		s = strconv.FormatUint(uint64(v), 10)
	case uint64:
		s = strconv.FormatUint(v, 10)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		s = strconv.FormatFloat(v, 'f', -1, 64)
	}

	if len(elems) == 1 {
		return s
	}

	n := len(sep) * (len(elems) - 1)
	for _, elem := range elems {
		switch v := any(elem).(type) {
		case int:
			n += len(strconv.Itoa(v))
		case int8:
			n += len(strconv.FormatInt(int64(v), 10))
		case int16:
			n += len(strconv.FormatInt(int64(v), 10))
		case int32:
			n += len(strconv.FormatInt(int64(v), 10))
		case int64:
			n += len(strconv.FormatInt(int64(v), 10))
		case uint:
			n += len(strconv.FormatUint(uint64(v), 10))
		case uint8:
			n += len(strconv.FormatUint(uint64(v), 10))
		case uint16:
			n += len(strconv.FormatUint(uint64(v), 10))
		case uint32:
			n += len(strconv.FormatUint(uint64(v), 10))
		case uint64:
			n += len(strconv.FormatUint(uint64(v), 10))
		case float32:
			n += len(strconv.FormatFloat(float64(v), 'f', -1, 32))
		case float64:
			n += len(strconv.FormatFloat(v, 'f', -1, 64))
		}
	}

	var b []byte
	b = make([]byte, n)
	bp := copy(b, s)
	for _, elem := range elems[1:] {
		bp += copy(b[bp:], sep)
		switch v := any(elem).(type) {
		case int:
			bp += copy(b[bp:], strconv.Itoa(v))
		case int8:
			bp += copy(b[bp:], strconv.FormatInt(int64(v), 10))
		case int16:
			bp += copy(b[bp:], strconv.FormatInt(int64(v), 10))
		case int32:
			bp += copy(b[bp:], strconv.FormatInt(int64(v), 10))
		case int64:
			bp += copy(b[bp:], strconv.FormatInt(v, 10))
		case uint:
			bp += copy(b[bp:], strconv.FormatUint(uint64(v), 10))
		case uint8:
			bp += copy(b[bp:], strconv.FormatUint(uint64(v), 10))
		case uint16:
			bp += copy(b[bp:], strconv.FormatUint(uint64(v), 10))
		case uint32:
			bp += copy(b[bp:], strconv.FormatUint(uint64(v), 10))
		case uint64:
			bp += copy(b[bp:], strconv.FormatUint(v, 10))
		case float32:
			bp += copy(b[bp:], strconv.FormatFloat(float64(v), 'f', -1, 32))
		case float64:
			bp += copy(b[bp:], strconv.FormatFloat(v, 'f', -1, 64))
		}
	}
	return string(b)
}

// SplitToNumbers splits a string into a slice of numeric type T using the given separator
func SplitToNumbers[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](s, sep string) ([]T, error) {
	if s == "" {
		return []T{}, nil
	}

	strParts := strings.Split(s, sep)
	result := make([]T, len(strParts))

	for i, part := range strParts {
		var err error
		var val any

		switch any(result[0]).(type) {
		case int:
			val, err = strconv.Atoi(part)
		case int8:
			v, e := strconv.ParseInt(part, 10, 8)
			val, err = int8(v), e
		case int16:
			v, e := strconv.ParseInt(part, 10, 16)
			val, err = int16(v), e
		case int32:
			v, e := strconv.ParseInt(part, 10, 32)
			val, err = int32(v), e
		case int64:
			val, err = strconv.ParseInt(part, 10, 64)
		case uint:
			v, e := strconv.ParseUint(part, 10, 0)
			val, err = uint(v), e
		case uint8:
			v, e := strconv.ParseUint(part, 10, 8)
			val, err = uint8(v), e
		case uint16:
			v, e := strconv.ParseUint(part, 10, 16)
			val, err = uint16(v), e
		case uint32:
			v, e := strconv.ParseUint(part, 10, 32)
			val, err = uint32(v), e
		case uint64:
			val, err = strconv.ParseUint(part, 10, 64)
		case float32:
			v, e := strconv.ParseFloat(part, 32)
			val, err = float32(v), e
		case float64:
			val, err = strconv.ParseFloat(part, 64)
		}

		if err != nil {
			return nil, err
		}
		result[i] = val.(T)
	}

	return result, nil
}
