package fingerprint

import (
	"testing"
	"fmt"
	"unsafe"
)

var data = "this is fingerprint test, test it hehehehehehehehehehehehehe.11sdkfjslllllllllllllllllllllllla"+
			"jeisfeksgnksgjskldmnkmnsgkljejsndkgnkejkslnkdnljo012i904i2orjqiwur90awfjlsajf93wjrfaeswuf93wjfi23us"+
			"sidogk2903ifk29uvjxdsvid90svjdogje902ujg34tj23928u55j2rmfoiew809sjavolaswie3r9023u5j2j320582303kfjw9e8vflsdejf9"

func TestFingerprint(t *testing.T){
	var fp FingerPrint
	fp.CalcFp([]byte(data))
	h1, h2 := uint64(0), uint64(0)
	h1 = *(*uint64)(unsafe.Pointer(&fp.data[0]))
	h2 = *(*uint64)(unsafe.Pointer(&fp.data[4]))
	fmt.Println("the data is: ", h1, h2)
}

func BenchmarkFingerPrint(b *testing.B){
	buf := make([]byte, 131072)
	for i := 0 ; i < 131072; i++{
		buf[i] = 'h'
	}

	b.ResetTimer()
	for i:=0; i < b.N; i++{
		var fp FingerPrint
		fp.CalcFp(buf)
	}
}
