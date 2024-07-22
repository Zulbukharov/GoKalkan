package ckalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include <strings.h>
// #include "KalkanCrypt.h"
//
// unsigned long hashData(char *algorithm, int flags, char *inData, int inDataLength, unsigned char *outData, int *outDataLength) {
//     bzero(outData, *outDataLength);
//     return kc_funcs->HashData(algorithm, flags, inData, inDataLength, outData, outDataLength);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

type HashAlgo string

const (
	HashAlgoSHA256 HashAlgo = "sha256"
	HashAlgoGOST95 HashAlgo = "Gost34311_95"
)

// HashData возвращается хеш dataB64 в base64.
func (cli *Client) HashData(algo HashAlgo, dataB64 string, flags Flag) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			if err != nil {
				err = fmt.Errorf("%w: panic: %s", err, r)
				return
			}

			err = fmt.Errorf("%w: %s", ErrPanic, r)
		}
	}()

	cli.mu.Lock()
	defer cli.mu.Unlock()

	kcAlgo := C.CString(string(algo))
	defer C.free(unsafe.Pointer(kcAlgo))

	kcInData := C.CString(dataB64)
	defer C.free(unsafe.Pointer(kcInData))
	inDataLength := len(dataB64)

	outDataLength := 50000 + 2*inDataLength
	outData := C.malloc(C.ulong(C.sizeof_uchar * outDataLength))
	defer C.free(outData)

	rc := int(C.hashData(
		kcAlgo,
		C.int(int(flags)),
		kcInData,
		C.int(inDataLength),
		(*C.uchar)(outData),
		(*C.int)(unsafe.Pointer(&outDataLength)),
	))

	return C.GoString((*C.char)(outData)), cli.wrapError(rc)
}
