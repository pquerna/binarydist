package snapdiff

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPatch(t *testing.T) {
	old := mustOpen("testdata/sample.old")
	defer old.Close()
	patch := mustOpen("testdata/sample.patch")
	defer patch.Close()

	tmp, err := ioutil.TempFile("", "snapdiff")
	require.NoError(t, err)
	defer tmp.Close()
	defer os.Remove(tmp.Name())

	err = Patch(old, tmp, patch)
	require.NoError(t, err)
	tmp.Close()

	expected, err := ioutil.ReadFile("testdata/sample.new")
	patched, err := ioutil.ReadFile(tmp.Name())
	require.NoError(t, err)
	require.Equal(t, expected, patched)
}

func TestPatchEmpty(t *testing.T) {
	old := mustOpen("testdata/sample.old")
	defer old.Close()

	tmp, err := ioutil.TempFile("", "snapdiff")
	require.NoError(t, err)
	defer tmp.Close()
	defer os.Remove(tmp.Name())

	err = Patch(old, tmp, bytes.NewReader([]byte{}))
	require.Error(t, err)
}
