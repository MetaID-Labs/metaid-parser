package decode

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/MetaID-Labs/metaid-parser/util"
)

func GetTxHash(rawTxByte []byte) string {
	limit := len(rawTxByte)
	if limit == 0 {
		return ""
	}
	index := 0
	if index+4 > limit {
		return ""
	}
	versionByte := rawTxByte[index : index+4]
	version := uint64(binary.LittleEndian.Uint32(versionByte))
	if version >= 10 {
		rawTxByte = GetTxNewHash(rawTxByte)
	}

	txHash := util.DoubleSHA256(rawTxByte)
	for i := 0; i < len(txHash)/2; i++ {
		h := txHash[len(txHash)-1-i]
		txHash[len(txHash)-1-i] = txHash[i]
		txHash[i] = h
	}
	return hex.EncodeToString(txHash)
}

func GetTxNewHash(rawTxByte []byte) []byte {
	var (
		newRawTxByte []byte
	)
	transaction, err := DecodeRawTransaction(rawTxByte)
	if err != nil {
		return newRawTxByte
	}
	newRawTxByte = getTxNewRawByte(transaction)
	return newRawTxByte
}

func getTxNewRawByte(transaction *RawTransaction) []byte {
	var (
		newRawTxByte   []byte
		newInputsByte  []byte
		newInputs2Byte []byte
		newOutputsByte []byte
	)
	newRawTxByte = append(newRawTxByte, transaction.Version...)
	newRawTxByte = append(newRawTxByte, transaction.LockTime...)
	newRawTxByte = append(newRawTxByte, util.Uint32ToLittleEndianBytes(uint32(transaction.inSize))...)
	newRawTxByte = append(newRawTxByte, util.Uint32ToLittleEndianBytes(uint32(transaction.outSize))...)

	for _, in := range transaction.Vins {
		newInputsByte = append(newInputsByte, in.TxID...)
		newInputsByte = append(newInputsByte, in.Vout...)
		newInputsByte = append(newInputsByte, in.sequence...)

		newInputs2Byte = append(newInputs2Byte, util.SHA256(in.scriptSig)...)
	}
	newRawTxByte = append(newRawTxByte, util.SHA256(newInputsByte)...)
	newRawTxByte = append(newRawTxByte, util.SHA256(newInputs2Byte)...)

	for _, out := range transaction.Vouts {
		newOutputsByte = append(newOutputsByte, out.amount...)
		newOutputsByte = append(newOutputsByte, util.SHA256(out.lockScript)...)
	}
	newRawTxByte = append(newRawTxByte, util.SHA256(newOutputsByte)...)
	return newRawTxByte
}
