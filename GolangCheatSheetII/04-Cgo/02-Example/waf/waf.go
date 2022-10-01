package waf

// #cgo CFLAGS: -I/usr/local/modsecurity/include
// #cgo LDFLAGS: -L/usr/local/modsecurity/lib/ -Wl,-rpath -Wl,/usr/local/modsecurity/lib/ -lmodsecurity
// #include "waf.h"
import "C"
import (
	"log"
	"time"
	"unsafe"
)

func InitModSec() {
	log.Printf("Init Mod Sec")
	C.InitModSec()
}

func ProcessModSec(url string) int {
	Curi := C.CString(url)
	defer C.free(unsafe.Pointer(Curi))
	start := time.Now()
	inter := int(C.ProcessModSec(Curi))
	elapsed := time.Since(start)
	log.Printf("\n========\nmodsec()=%i, elapsed: %s", inter, elapsed)
	return inter
}
