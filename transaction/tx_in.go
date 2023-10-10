package decode

import (
	"encoding/hex"
	"github.com/MetaID-Labs/metaid-parser/util"
)

type TxIn struct {
	inType     int
	TxID       []byte
	Vout       []byte
	scriptSig  []byte
	sequence   []byte
	lockScript []byte
}

const (
	TypeEmpty    = 0
	TypeP2PK     = 1
	TypeP2PKH    = 2
	TypeP2WPKH   = 3
	TypeBech32   = 4
	TypeMultiSig = 5
)

const (
	Type_Empty    = "Empty"
	Type_P2PK     = "P2PK"
	Type_P2PKH    = "P2PKH"
	Type_P2PSH    = "P2PSH"
	Type_P2WPKH   = "P2WPKH"
	Type_Bech32   = "Bech32"
	Type_MultiSig = "MultiSig"
	Type_NullData = "NullData"
)

func (in TxIn) GetUTXOType() int {
	return in.inType
}

func (in TxIn) GetUTXOTypeStr() string {
	switch in.inType {
	case TypeEmpty:
		return Type_Empty
	case TypeP2PK:
		return Type_P2PK
	case TypeP2PKH:
		return Type_P2PKH
	}
	return ""
}

func (in TxIn) GetTxID() string {
	return util.ReverseBytesToHex(in.TxID)
}

func (in TxIn) GetVout() uint32 {
	return util.LittleEndianBytesToUint32(in.Vout)
}

func (in TxIn) GetScriptSig() string {
	return hex.EncodeToString(in.scriptSig)
}

func (in TxIn) GetSequence() uint32 {
	return util.LittleEndianBytesToUint32(in.sequence)
}
