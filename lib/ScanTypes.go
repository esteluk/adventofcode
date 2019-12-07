package lib

import "unicode/utf8"

func isComma(r rune) bool {
	return ',' == r
}

// ScanCommaSeparated returns comma separated values
func ScanCommaSeparated(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0

	// Scan until comma, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if isComma(r) {
			return i + width, data[start:i], nil
		}

	}

	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}

	// Request more data.
	return start, nil, nil
}
