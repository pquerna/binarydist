// Package snapdiff implements binary diff and patch as described on
// http://www.daemonology.net/bsdiff/, but changed to use snappy as the compression
package snapdiff

var magic = [8]byte{'S', 'N', 'D', 'I', 'F', '4', '0'}

// File format:
//   0       8    "BSDIFF40"
//   8       8    X
//   16      8    Y
//   24      8    sizeof(newfile)
//   32      X    bzip2(control block)
//   32+X    Y    bzip2(diff block)
//   32+X+Y  ???  bzip2(extra block)
// with control block a set of triples (x,y,z) meaning "add x bytes
// from oldfile to x bytes from the diff block; copy y bytes from the
// extra block; seek forwards in oldfile by z bytes".
type header struct {
	Magic   [8]byte
	CtrlLen int64
	DiffLen int64
	NewSize int64
}
