// +build ppc

package bolt

// maxMapSize represents the largest mmap size supported by
const maxMapSize = 0x7FFFFFFF // 2GB

// maxAllocSize is the size used when creating array pointers.
const maxAllocSize = 0xFFFFFFF
