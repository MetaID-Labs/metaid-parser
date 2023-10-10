package decode

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"github.com/MetaID-Labs/metaid-parser/util"
)

type RawTransaction struct {
	TxID          string
	Size          uint64
	Hex           string
	BlockHash     string
	BlockHeight   uint64
	Confirmations uint64
	Blocktime     int64
	inSize        uint64
	outSize       uint64

	Version  []byte
	Vins     []TxIn
	Vouts    []TxOut
	LockTime []byte
	Witness  bool
}

func DecodeRawTransaction(txBytes []byte) (*RawTransaction, error) {
	limit := len(txBytes)
	if limit == 0 {
		return nil, errors.New("Invalid transaction data!")
	}
	var rawTx RawTransaction
	index := 0
	if index+4 > limit {
		return nil, errors.New("Invalid transaction data length!")
	}
	rawTx.Version = txBytes[index : index+4]
	index += 4

	if index+2 > limit {
		return nil, errors.New("Invalid transaction data length!")
	}

	//if txBytes[index] == SegWitSymbol {
	//	if txBytes[index+1] != SegWitVersion {
	//		return nil, errors.New("Invalid witness symbol!")
	//	}
	//	rawTx.Witness = true
	//	index += 2
	//}

	if index+1 > limit {
		return nil, errors.New("Invalid transaction data length!")
	}

	icount, lenth := util.DecodeVarIntForTx(txBytes[index : index+9])
	numOfVins := icount
	rawTx.inSize = uint64(numOfVins)
	index += lenth

	if numOfVins == 0 {
		return nil, errors.New("Invalid transaction data!")
	}
	for i := 0; i < numOfVins; i++ {
		var tmpTxIn TxIn

		if index+32 > limit {
			return nil, errors.New("Invalid transaction data length!")
		}
		tmpTxIn.TxID = txBytes[index : index+32]
		index += 32

		if index+4 > limit {
			return nil, errors.New("Invalid transaction data length!")
		}
		tmpTxIn.Vout = txBytes[index : index+4]
		index += 4

		if index+1 > limit {
			return nil, errors.New("Invalid transaction data length!")
		}

		vnumber := txBytes[index : index+9]
		icount, lenth = util.DecodeVarIntForTx(vnumber)
		scriptLen := icount
		index += lenth

		tmpTxIn.scriptSig = txBytes[index : index+scriptLen]
		index += scriptLen

		tmpTxIn.sequence = txBytes[index : index+4]
		index += 4
		rawTx.Vins = append(rawTx.Vins, tmpTxIn)
	}

	if index+1 > limit {
		return nil, errors.New("Invalid transaction data length!")
	}

	icount, lenth = util.DecodeVarIntForTx(txBytes[index : index+9])
	numOfVouts := icount
	rawTx.outSize = uint64(numOfVouts)
	index += lenth

	if numOfVouts == 0 {
		return nil, errors.New("Invalid transaction data!")
	}

	for i := 0; i < numOfVouts; i++ {
		var tmpTxOut TxOut
		tmpTxOut.n = uint((i))
		if index+8 > limit {
			return nil, errors.New("Invalid transaction data length!")
		}
		tmpTxOut.amount = txBytes[index : index+8]
		index += 8

		if index+1 > limit {
			return nil, errors.New("Invalid transaction data length!")
		}

		vnumber := txBytes[index : index+9]
		icount, lenth = util.DecodeVarIntForTx(vnumber)
		lockScriptLen := icount
		index += lenth

		if lockScriptLen == 0 {
			return nil, errors.New("Invalid transaction data!")
		}
		if index+int(lockScriptLen) > limit {
			return nil, errors.New("Invalid transaction data length!")
		}
		tmpTxOut.lockScript = txBytes[index : index+int(lockScriptLen)]
		index += int(lockScriptLen)
		rawTx.Vouts = append(rawTx.Vouts, tmpTxOut)
	}

	if index+4 > limit {
		return nil, errors.New("Invalid transaction data length!")
	}
	rawTx.LockTime = txBytes[index : index+4]
	index += 4

	if index != limit {
		return nil, errors.New("Too much transaction data!")
	}

	if uint64(binary.LittleEndian.Uint32(rawTx.Version)) < 10 {
		rawTx.TxID = util.GetTxID(hex.EncodeToString(txBytes))
	} else {
		newRawTxByte := getTxNewRawByte(&rawTx)
		rawTx.TxID = util.GetTxID(hex.EncodeToString(newRawTxByte))
	}

	return &rawTx, nil
}
