package httputil

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

// @Author Chen Zikang
// @Email  parakovo@gmail.com
// @Since  2022-09-09

// Redirect to url by setting response header.
// Set location to url and http status to code.
// Shouldn't redirect for POST or HEAD request.
// Additionally, Code should be 3xx.
func Redirect(w http.ResponseWriter, url string, code int) {
	h := w.Header()

	// RFC 7231 notes that a short HTML body is usually included in
	// the response because older user agents may not understand 301/307.
	// Do it only if the request didn't already have a Content-Type header.
	h.Set("Location", hexEscapeNonASCII(url))
	h.Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)

	// Shouldn't send the body for POST or HEAD; that leaves GET.
	body := "<a href=\"" + htmlEscape(url) + "\">" + http.StatusText(code) + "</a>.\n"
	fmt.Fprintln(w, body)
}

func hexEscapeNonASCII(s string) string {
	newLen := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= utf8.RuneSelf {
			newLen += 3
		} else {
			newLen++
		}
	}
	if newLen == len(s) {
		return s
	}
	b := make([]byte, 0, newLen)
	for i := 0; i < len(s); i++ {
		if s[i] >= utf8.RuneSelf {
			b = append(b, '%')
			b = strconv.AppendInt(b, int64(s[i]), 16)
		} else {
			b = append(b, s[i])
		}
	}
	return string(b)
}

var htmlReplacer = strings.NewReplacer(
	"&", "&amp;",
	"<", "&lt;",
	">", "&gt;",
	// "&#34;" is shorter than "&quot;".
	`"`, "&#34;",
	// "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
	"'", "&#39;",
)

func htmlEscape(s string) string {
	return htmlReplacer.Replace(s)
}
