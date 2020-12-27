package simplesurt

import (
	"fmt"
	"net/url"
	"strings"
)

// Format canoncalizes a given URI string and returns it as a formatted SURT.
// A URI <scheme://domain.tld/path?query> has SURT form
// <scheme://(tld,domain,)/path?query>
// For more information see the manual at:
// http://crawler.archive.org/articles/user_manual/glossary.html#surt
func Format(uri string) (string, error) {
	// canoncalize the uri
	canon, err := Canonicalize(uri)
	if err != nil {
		return "", err
	}
	parsed, err := url.Parse(canon)
	if err != nil {
		return "", err
	}
	// build the SURT
	domains := strings.Split(parsed.Hostname(), ".")
	domains = reverse(domains)
	surt := fmt.Sprintf("%s://", parsed.Scheme)
	surt = fmt.Sprintf("%s(%s,", surt, strings.Join(domains, ","))
	if port := parsed.Port(); port != "" {
		surt = fmt.Sprintf("%s:%s", surt, port)
	}
	surt = fmt.Sprintf("%s)", surt)
	if path := parsed.Path; path != "" {
		surt = fmt.Sprintf("%s%s", surt, path)
	}
	if query := parsed.Query().Encode(); query != "" {
		surt = fmt.Sprintf("%s?%s", surt, query)
	}
	return surt, nil
}

// reverse reverses the order of the given string array
func reverse(ss []string) []string {
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}
	return ss
}
