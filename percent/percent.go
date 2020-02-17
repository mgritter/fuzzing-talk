package percent

func PercentEncode(input string) (string, error) {
	result := make([]byte, 0, len(input))
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '!':
			result = append(result, []byte("%21")...)
		case '#':
			result = append(result, []byte("%23")...)
		case '$':
			result = append(result, []byte("%24")...)
		case '&':
			result = append(result, []byte("%26")...)
		case '\'':
			result = append(result, []byte("%27")...)
		case '(':
			result = append(result, []byte("%28")...)
		case ')':
			result = append(result, []byte("%29")...)
		case '*':
			result = append(result, []byte("%2a")...)
		case '+':
			result = append(result, []byte("%2B")...)
		case ',':
			result = append(result, []byte("%2C")...)
		case '/':
			result = append(result, []byte("%2F")...)
		case ':':
			result = append(result, []byte("%3A")...)
		case ';':
			result = append(result, []byte("%3B")...)
		case '=':
			result = append(result, []byte("%3D")...)
		case '?':
			result = append(result, []byte("%3F")...)
		case '@':
			result = append(result, []byte("%40")...)
		case '[':
			result = append(result, []byte("%5B")...)
		case ']':
			result = append(result, []byte("%5D")...)
		default:
			result = append(result, input[i])
		}
	}
	return string(result), nil
}
