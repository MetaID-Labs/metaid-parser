package decode

import (
	"encoding/hex"
	"fmt"
	"github.com/MetaID-Labs/metaid-parser/parser"
	"github.com/MetaID-Labs/metaid-parser/util"
	_ "strconv"
	"testing"
)

func Test_TxHex2Transaction(t *testing.T) {
	txHex := "0a0000000147b83d05df8c946bd8b435019dd503fe98f8bb4e90585bf77759eb23094f910c010000006a4730440220066f947a78d70d4afb08699b81ead3021ee3d0bbfbb57cf1b8eb95457102100e02205ea18fc1a7fd041a73da1a7972dd499b794589b4f7489770c6c98486e296755c412102e45470a184e84658ea7aa412470c674cf8fa19afd9802e40a656f44c8eb95601ffffffff02000000000000000074006a046d65746142303265343534373061313834653834363538656137616134313234373063363734636638666131396166643938303265343061363536663434633865623935363031044e554c4c066d657461696404526f6f74044e554c4c044e554c4c044e554c4c044e554c4c044e554c4c46030300000000001976a91412e05e38481043bb00cb78be5f190940567d0cab88ac00000000"

	txBytes, err := hex.DecodeString(txHex)
	if err != nil {
		fmt.Println("Invalid transaction hex string!")
		fmt.Println("Err:", err)
		return
	}

	transaction, err := DecodeRawTransaction(txBytes)
	if err != nil {
		fmt.Println("Err:", err)
		return
	}

	fmt.Println("Transaction:")
	fmt.Println("Version:", util.LittleEndianBytesToUint32(transaction.Version))
	fmt.Println("LockTime:", util.LittleEndianBytesToUint32(transaction.LockTime))
	fmt.Println("VinSize:", int64(transaction.inSize))
	fmt.Println("Vin:[")
	for _, v := range transaction.Vins {
		fmt.Println("	{")
		fmt.Println("		TxId:", v.GetTxID())
		fmt.Println("		Index:", v.GetVout())
		fmt.Println("		Type:", v.GetUTXOTypeStr())
		fmt.Println("		ScriptSig:", v.GetScriptSig())
		//fmt.Println("		ScriptPub:", v.GetScriptPub()))
		//fmt.Println("		ScriptMulti:", v.GetScriptMulti()))
		fmt.Println("		Sequence:", v.GetSequence())
		fmt.Println("	}")
	}
	fmt.Println("]")

	fmt.Println("VoutSize:", int64(transaction.outSize))
	fmt.Println("Vout:[")
	for _, v := range transaction.Vouts {
		fmt.Println("	{")
		fmt.Println("		Amount:", v.GetAmount())
		fmt.Println("		n:", v.GetIndex())
		fmt.Println("		LockScript:", v.GetLockScript())

		sType, pkh, opStrs, parts, err := parser.DecodeCommonOutScript(v.GetLockScript())
		if err != nil {
			fmt.Println("DecodeCommonOutScript err:", err)
			return
		}
		switch sType {
		case parser.TypeP2PK:
			fmt.Println("		sType:", Type_P2PK)
			fmt.Println("		pkh:", pkh)
			fmt.Println("		address:", util.PublicKeyHashToAddress(pkh))
			break
		case parser.TypeP2PKH:
			fmt.Println("		sType:", Type_P2PKH)
			fmt.Println("		pkh:", pkh)
			fmt.Println("		address:", util.PublicKeyHashToAddress(pkh))
			break
		}
		fmt.Println("		opStrs:", opStrs)
		fmt.Println("		parts:", parts)
		fmt.Println("	}")
	}
	fmt.Println("]")

	// 0x14 0x00
}
