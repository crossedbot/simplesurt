package simplesurt

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	// Regular expressions
	controlRe  = regexp.MustCompile(`[\x00-\x1F\x7F]`)
	nonAsciiRe = regexp.MustCompile(`[^[:ascii:]]`)
)

// normalize normalizes a given URI string or fragment of a string
func normalize(s string) string {
	s = strings.ToLower(s)
	s = controlRe.ReplaceAllString(s, "")
	s = nonAsciiRe.ReplaceAllString(s, "")
	return s
}

// unquote unquotes a (multi-)quoted string
func unquote(v string) string {
	for {
		next, err := strconv.Unquote(v)
		if err != nil {
			next = v
		}
		if next == v {
			return next
		}
		v = next
	}
}

// escape unquotes and quotes a given string to ASCII
func escape(v string) string {
	return strings.Trim(strconv.QuoteToASCII(unquote(v)), `"`)
}

// buildUrl returns a URL string for the given parts
func buildUrl(scheme, hostname, port, path, query string) string {
	uri := fmt.Sprintf("%s://%s", scheme, hostname)
	if port != "" {
		uri = fmt.Sprintf("%s:%s", uri, port)
	}
	if path != "" {
		if !strings.HasPrefix(path, "/") {
			uri = fmt.Sprintf("%s/%s", uri, path)
		} else {
			uri = fmt.Sprintf("%s%s", uri, path)
		}
	}
	if query != "" {
		if !strings.HasPrefix(query, "?") {
			uri = fmt.Sprintf("%s?%s", uri, query)
		} else {
			uri = fmt.Sprintf("%s%s", uri, query)
		}
	}
	return uri
}
