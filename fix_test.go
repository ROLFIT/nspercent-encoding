// unescape_test
package NSPercentEncoding

import (
	"net/url"
	"testing"
)

func testFix(t *testing.T, escaped, unescaped string) {
	if r, err := url.QueryUnescape(FixNonStandardPercentEncoding(escaped)); (r != unescaped) || (err != nil) {
		t.Errorf("String should be \"%s\", was \"%s\" Error: %s", unescaped, r, err)
	}
}

func TestFix(t *testing.T) {
	var data = []struct {
		escaped   string
		unescaped string
	}{
		{"%D0%9F%D1%80%D0%B8%D0%B2%D0%B5%D1%82%2C+%D0%BC%D0%B8%D1%80+%2B+%D0%AF%21%21%21", "Привет, мир + Я!!!"},
		{"%25%D0%94%D0%B0%D0%B2%D1%8B%D0%B4%D0%BE%D0%B2%25%D0%98%D0%B2%D0%B0%D0%BD%25", "%Давыдов%Иван%"},
		{"%25%u0414%u0430%u0432%u044B%u0434%u043E%u0432%25%u0418%u0432%u0430%u043D%25", "%Давыдов%Иван%"},
		{"%25%25", "%%"},
	}
	for _, d := range data {

		testFix(t, d.escaped, d.unescaped)
	}
}
