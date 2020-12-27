package simplesurt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanonicalize(t *testing.T) {
	expected := "https://" +
		"example.com:8443" +
		"/hello/world" +
		"?empty=&one=two+&three=four"
	actual, err := Canonicalize("HTTPS\t://" +
		"www.EXAMPLE..COM:8443" +
		"/hello///rm/../world/./" +
		"?ONE=two &empty=&Three=four",
	)
	require.Nil(t, err)
	require.Equal(t, expected, actual)
}

func TestCanonicalizeScheme(t *testing.T) {
	expected := "https"
	actual := CanonicalizeScheme("HTTPS\t")
	require.Equal(t, expected, actual)
}

func TestCanonicalizeHostname(t *testing.T) {
	expected := "example.com"
	actual := CanonicalizeHostname("www.EXAMPLE..COM")
	require.Equal(t, expected, actual)
}

func TestCanonicalizePort(t *testing.T) {
	expected := ""
	actual := CanonicalizePort("http", "80")
	require.Equal(t, expected, actual)

	expected = "8080"
	actual = CanonicalizePort("http", "8080")
	require.Equal(t, expected, actual)
}

func TestCanonicalizePath(t *testing.T) {
	expected := "/hello/world"
	actual := CanonicalizePath("/hello///rm/../world/./")
	require.Equal(t, expected, actual)
}

func TestCanonicalizeQuery(t *testing.T) {
	expected := "empty=&one=two+&three=four"
	actual := CanonicalizeQuery("ONE=two &empty=&Three=four")
	require.Equal(t, expected, actual)
}
