package decode

import (
	"encoding/hex"
	"github.com/MetaID-Labs/metaid-parser/parser"
	"github.com/MetaID-Labs/metaid-parser/util"
)

type TxOut struct {
	n          uint
	amount     []byte
	lockScript []byte
	scriptType int64 //TODO
}

func (out TxOut) GetAmount() uint64 {
	return util.LittleEndianBytesToUint64(out.amount)
}

func (out TxOut) GetIndex() uint64 {
	return uint64(out.n)
}

func (out TxOut) GetLockScript() string {
	return hex.EncodeToString(out.lockScript)
	//TODO
}

func (out TxOut) decodeLockScript() string {
	_, _, opStrs, _, _ := parser.DecodeCommonOutScript(hex.EncodeToString(out.lockScript))
	return opStrs
}
