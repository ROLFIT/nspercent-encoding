// unescape
package NSPercentEncoding

import (
	"net/url"
	"unicode/utf8"
)

func ishex(c byte) bool {
	switch {
	case '0' <= c && c <= '9':
		return true
	case 'a' <= c && c <= 'f':
		return true
	case 'A' <= c && c <= 'F':
		return true
	}
	return false
}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

// unescape unescapes a string; the mode specifies
// which section of the URL string is being unescaped.
func UnEscape(s string) (string, error) {
	// Count %, check that they're well-formed.
	n := 0
	hasPlus := false
	hasEscaped := false
	for i := 0; i < len(s); {
		switch s[i] {
		case '%':
			hasEscaped = true
			if i+1 < len(s) && s[i+1] == 'u' {
				if i+4 >= len(s) || !ishex(s[i+2]) || !ishex(s[i+3]) || !ishex(s[i+4]) || !ishex(s[i+5]) {
					s = s[i:]
					if len(s) > 5 {
						s = s[0:5]
					}
					return "", url.EscapeError(s)
				}
				i += 6
				n += 2
			} else {
				if i+2 >= len(s) || !ishex(s[i+1]) || !ishex(s[i+2]) {
					s = s[i:]
					if len(s) > 3 {
						s = s[0:3]
					}
					return "", url.EscapeError(s)
				}
				i += 3
				n++
			}
		case '+':
			hasPlus = true
			i++
			n++
		default:
			i++
			n++
		}
	}

	if !hasEscaped && !hasPlus {
		return s, nil
	}
	t := make([]byte, n)
	j := 0
	for i := 0; i < len(s); {
		switch s[i] {
		case '%':
			if s[i+1] == 'u' {
				var v rune
				v = v<<4 | rune(unhex(s[i+2]))
				v = v<<4 | rune(unhex(s[i+3]))
				v = v<<4 | rune(unhex(s[i+4]))
				v = v<<4 | rune(unhex(s[i+5]))

				b := make([]byte, utf8.RuneLen(v))
				utf8.EncodeRune(b, v)
				for _, c := range b {
					t[j] = c
					j++
				}
				i += 6
			} else {
				t[j] = unhex(s[i+1])<<4 | unhex(s[i+2])
				j++
				i += 3
			}
		case '+':
			t[j] = ' '
			j++
			i++
		default:
			t[j] = s[i]
			j++
			i++
		}
	}
	return string(t), nil
}
