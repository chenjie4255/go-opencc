package opencc

// #cgo LDFLAGS: -lopencc
/*
#include<opencc/opencc.h>
*/
import "C"

type Config struct {
	cfFile C.opencc_t
}

type ConfigType int

const (
	ConfigTypeS2T  ConfigType = 1
	ConfigTypeS2TW ConfigType = 2
)

func NewConfig(t ConfigType) *Config {
	cf := Config{}
	if t == ConfigTypeS2T {
		cf.cfFile = C.opencc_open(C.CString("s2t.json"))
	} else if t == ConfigTypeS2TW {
		cf.cfFile = C.opencc_open(C.CString("s2tw.json"))
	}

	return &cf
}

func (c *Config) Close() {
	if c.cfFile != nil {
		C.opencc_close(c.cfFile)
	}
}

func Tr(str string, config *Config) string {
	ccResult := C.opencc_convert_utf8(config.cfFile, C.CString(str), C.size_t(len(str)))
	if ccResult == nil {
		return ""
	}
	result := C.GoString(ccResult)
	C.opencc_convert_utf8_free(ccResult)
	return result
}
