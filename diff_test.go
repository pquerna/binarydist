package snapdiff

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiff(t *testing.T) {
	old := mustOpen("testdata/sample.old")
	defer old.Close()
	new := mustOpen("testdata/sample.new")
	defer new.Close()

	buf := &bytes.Buffer{}
	err := Diff(old, new, buf)
	require.NoError(t, err)

	sample, err := ioutil.ReadFile("testdata/sample.patch")
	require.NoError(t, err)
	require.Equal(t, sample, buf.Bytes())
}
