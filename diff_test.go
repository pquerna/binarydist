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

func TestDiffPatchSimpleStrings(t *testing.T) {
	patch := &bytes.Buffer{}

	old := bytes.NewReader([]byte("hello world"))
	new := bytes.NewBufferString("hello snappy")

	err := Diff(old, new, patch)
	require.NoError(t, err)

	next := &bytes.Buffer{}
	old.Seek(0, 0)
	err = Patch(old, next, bytes.NewReader(patch.Bytes()))
	require.NoError(t, err)
	require.Equal(t, []byte("hello snappy"), next.Bytes())
}
