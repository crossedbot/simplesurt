package simplesurt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNormalize(t *testing.T) {
	v := "Hellο\tWοrld" // o's are actually omicrons
	expected := "hellwrld"
	actual := normalize(v)
	require.Equal(t, expected, actual)
}

func TestUnquote(t *testing.T) {
	expected := "\u2639\u2639"
	actual := unquote("`\"\u2639\u2639\"`")
	require.Equal(t, expected, actual)
}

func TestEscape(t *testing.T) {
	v := "`\"Hellο\tWοrld\"`"
	expected := "Hell\\u03bf\\tW\\u03bfrld"
	actual := escape(v)
	require.Equal(t, expected, actual)
}

func TestBuildUrl(t *testing.T) {
	expected := "https://example.com:8443/hello/world?one=two"
	actual := buildUrl(
		"https",       // scheme
		"example.com", // hostname
		"8443",        // port
		"hello/world", // path
		"one=two",     // query
	)
	require.Equal(t, expected, actual)
}
