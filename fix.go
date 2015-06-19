// unescape
package NSPercentEncoding

import (
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

func FixNonStandardPercentEncoding(s string) string {
	r := make([]byte, len(s))
	for i := 0; i < len(s); {
		switch s[i] {
		case '%':
			if i+1 < len(s) && s[i+1] == 'u' {
				t := s[i+2:]

				if len(t) >= 4 && ishex(t[0]) && ishex(t[1]) && ishex(t[2]) && ishex(t[3]) {
					var v rune
					v = v<<4 | rune(unhex(t[0]))
					v = v<<4 | rune(unhex(t[1]))
					v = v<<4 | rune(unhex(t[2]))
					v = v<<4 | rune(unhex(t[3]))

					b := make([]byte, utf8.RuneLen(v))
					if utf8.EncodeRune(b, v) != 2 {
						// Оставляем без изменений. Пусть разбирается url.QueryUnescape
						r[i] = s[i]
						i++
					} else {
						r[i] = '%'
						r[i+1] = "0123456789ABCDEF"[b[0]>>4]
						r[i+2] = "0123456789ABCDEF"[b[0]&15]
						r[i+3] = '%'
						r[i+4] = "0123456789ABCDEF"[b[1]>>4]
						r[i+5] = "0123456789ABCDEF"[b[1]&15]
						i += 6
					}
				} else {
					r[i] = s[i]
					i++
				}

			} else {
				r[i] = s[i]
				i++
			}
		default:
			r[i] = s[i]
			i++
		}
	}
	return string(r)
}
