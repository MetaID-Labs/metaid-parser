package parser

import "strings"

const (
	Op0      = 0x00
	OpReturn = 0x6A
)

const (
	Op0Str      = "OP_0"
	OpReturnStr = "OP_RETURN"
)

var (
	OpMeta             = []byte{0x6d, 0x65, 0x74, 0x61}
	OpMVC              = []byte{0x6d, 0x76, 0x63}
	OP_CHAIN_FLAG_LIST = [][]byte{
		OpMeta,
		OpMVC,
	}
)

func IsChainFlag(opStr string) (bool, string) {
	for _, v := range OP_CHAIN_FLAG_LIST {
		if opStr == string(v) {
			return true, string(v)
		}
	}
	return false, ""
}

func CheckChainFlag(chainFlag string) bool {
	for _, v := range OP_CHAIN_FLAG_LIST {
		if strings.ToLower(chainFlag) == string(v) {
			return true
		}
	}
	return false
}
