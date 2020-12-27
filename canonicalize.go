package simplesurt

import (
	"net/url"
	"regexp"
	"strings"
)

var (
	// Regular Expressions
	dotsRe      = regexp.MustCompile(`\.\.+`)
	slashRe     = regexp.MustCompile(`\/\/+`)
	wwwDomainRe = regexp.MustCompile(`^www\d*$`)
)

// defaultSchemePorts are is map of common URI schemes are the ports commonly
// associated with them
var defaultSchemePorts = map[string]string{
	"http":   "80",
	"https":  "443",
	"ftp":    "21",
	"ssh":    "22",
	"telnet": "23",
}

// Canonicalize canonicalizes a given URI string
func Canonicalize(uri string) (string, error) {
	// normalize the URI and parse it
	uri = normalize(uri)
	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	// canonicalize each part separately
	scheme := CanonicalizeScheme(u.Scheme)
	hostname := CanonicalizeHostname(u.Hostname())
	port := CanonicalizePort(scheme, u.Port())
	path := CanonicalizePath(u.Path)
	query := CanonicalizeQuery(u.RawQuery)
	// rebuild the URI string
	u, err = url.Parse(buildUrl(scheme, hostname, port, path, query))
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

// CanonicalizeScheme canonicalizes the scheme part of a URI
func CanonicalizeScheme(scheme string) string {
	return normalize(scheme)
}

// CanonicalizeHostname canonicalizes the hostname part of a URI
func CanonicalizeHostname(hostname string) string {
	hostname = normalize(hostname)
	hostname = dotsRe.ReplaceAllString(hostname, ".")
	domains := strings.Split(hostname, ".")
	if len(domains) > 0 && wwwDomainRe.MatchString(domains[0]) {
		hostname = strings.Join(domains[1:], ".")
	}
	return escape(hostname)
}

// CanonicalizePort canonicalizes the port part of a URI
func CanonicalizePort(scheme, port string) string {
	if defaultSchemePorts[scheme] == port {
		return ""
	}
	return port
}

// CanonicalizePath canonicalizes the path part of a URI; a leading forward
// slash is expected and  trailing forward slashes are removed
func CanonicalizePath(path string) string {
	tracked := []string{}
	// normalize the path
	path = normalize(path)
	path = slashRe.ReplaceAllString(path, "/")
	// for each segment of the path remove dot segments and empty strings;
	// moving up one directory level on ".." segments
	parts := strings.Split(path, "/")
	for _, part := range parts[1:] {
		if part == "." || part == "" {
			continue
		} else if part == ".." {
			if len(tracked) > 0 {
				tracked = tracked[:len(tracked)-1]
			} else {
				tracked = append(tracked, part)
			}
		} else {
			tracked = append(tracked, part)
		}
	}
	tracked = append([]string{""}, tracked...)
	return escape(strings.Join(tracked, "/"))
}

// CanonicalizeQuery canonicalizes the search part of a URI
func CanonicalizeQuery(query string) string {
	query = normalize(query)
	m, _ := url.ParseQuery(query)
	return escape(m.Encode())
}
