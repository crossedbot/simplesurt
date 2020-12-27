package simplesurt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormat(t *testing.T) {
	expected := "https://" +
		"(com,example,:8443)" +
		"/hello/world" +
		"?empty=&one=two+&three=four"
	actual, err := Format("HTTPS\t://" +
		"www.EXAMPLE..COM:8443" +
		"/hello///rm/../world/./" +
		"?ONE=two &empty=&Three=four",
	)
	require.Nil(t, err)
	require.Equal(t, expected, actual)
}

func TestReverse(t *testing.T) {
	arr := []string{"one", "two", "three"}
	expected := []string{"three", "two", "one"}
	actual := reverse(arr)
	require.Equal(t, expected, actual)
}
