package fingerprint

import (
	"github.com/spaolacci/murmur3"
	"unsafe"
)
const fp_seed = uint32(0)

const FP_len = 16

type FingerPrint struct{
	data [FP_len]byte
}

func (fp *FingerPrint)CalcFp(data []byte) {
	h128 := murmur3.New128WithSeed(0)
	h128.Write(data)
	h1, h2 := h128.Sum128()
	*(*uint64)(unsafe.Pointer(&fp.data[0])) = h1
	*(*uint64)(unsafe.Pointer(&fp.data[4])) = h2
}


