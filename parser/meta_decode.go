package parser

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/MetaID-Labs/metaid-parser/util"
	"log"
)

/*
DecodeMetaOutScript :

	Used to parse script data in Outputs for parser data

return:

	PartsString - []string - The parts of script
	IsMeta      - bool     - Whether Meta
	Error       - error    - The message of error
*/
func DecodeMetaOutScript(scriptPubKey string, txId string) ([]string, bool, error) {
	//defer CatchPanic("DecodeMetaOutScript", txId)
	parts := make([]string, 0)
	scriptBytes, err := hex.DecodeString(scriptPubKey)
	if err != nil {
		log.Println("Invalid script hex string! Err:", err)
		return nil, false, err
	}

	limit := len(scriptBytes)
	if limit == 0 {
		return nil, false, errors.New("Invalid script data!")
	}
	index := 0

	//OP_0 OP_RETURN
	if index+2 > limit {
		return nil, false, errors.New("Invalid script data length!")
	}
	if scriptBytes[index] != Op0 && scriptBytes[index+1] != OpReturn {
		return nil, false, errors.New("Not a opReturn script.")
	}
	parts = append(parts, Op0Str, OpReturnStr)
	index += 2

	//ChainFlag
	if index+2 > limit {
		return nil, false, errors.New("Invalid script data length!")
	}
	icount, lenth := util.DecodeVarIntForScript(scriptBytes[index : index+9])
	if icount == -1 {
		fmt.Println("******************maybe op_Code - ", limit)
		return nil, false, errors.New("Invalid script data length! maybe op_Code")
	}
	index += lenth
	metaLen := icount
	// TODO +len之后 catch fix
	if index+metaLen > limit {
		return nil, false, errors.New("Invalid script data length!")
	}

	isChainFlag, chainFlagStr := IsChainFlag(string(scriptBytes[index : index+metaLen]))
	if !isChainFlag {
		return nil, false, errors.New("Not a metaid script.")
	}
	parts = append(parts, chainFlagStr)
	index += metaLen

	//address/publicKey
	if index+2 > limit {
		return nil, true, errors.New("Invalid metaid script data length!")
	}
	icount, lenth = util.DecodeVarIntForScript(scriptBytes[index : index+9])
	if icount == -1 {
		fmt.Println("******************maybe op_Code - ", limit)
		return nil, false, errors.New("Invalid script data length! maybe op_Code")
	}
	index += lenth
	addressLen := icount
	address := string(scriptBytes[index : index+addressLen])
	parts = append(parts, address)
	index += addressLen

	//parentTxId
	if index+2 > limit {
		return nil, true, errors.New("Invalid metaid script data length!")
	}
	icount, lenth = util.DecodeVarIntForScript(scriptBytes[index : index+9])
	if icount == -1 {
		fmt.Println("******************maybe op_Code - ", limit)
		return nil, false, errors.New("Invalid script data length! maybe op_Code")
	}
	index += lenth
	parentTxIdLen := icount
	parentTxId := string(scriptBytes[index : index+parentTxIdLen])
	parts = append(parts, parentTxId)
	index += parentTxIdLen

	//data content
	if index == limit {
		return parts, true, nil
	}

	completed := true
	for completed {
		icount, lenth = util.DecodeVarIntForScript(scriptBytes[index : index+9])
		if icount == -1 {
			fmt.Println("******************maybe op_Code - ", limit)
			return nil, false, errors.New("Invalid script data length! maybe op_Code")
		}
		index += lenth
		partLen := icount
		part := string(scriptBytes[index : index+partLen])
		parts = append(parts, part)
		index += partLen
		if index == limit {
			completed = false
		}
	}

	return parts, true, nil
}

// for metafile
func DecodeMetaOutScriptForMetaFile(script, encoding string) ([]string, bool, error) {
	parts := make([]string, 0)
	scriptBytes, err := hex.DecodeString(script)
	if err != nil {
		log.Println("Invalid script hex string! Err:", err)
		return nil, false, err
	}

	limit := len(scriptBytes)
	if limit == 0 {
		return nil, false, errors.New("Invalid script data!")
	}
	index := 0

	//OP_0 OP_RETURN
	if index+2 > limit {
		return nil, false, errors.New("Invalid script data length!")
	}
	if scriptBytes[index] != Op0 && scriptBytes[index+1] != OpReturn {
		return nil, false, errors.New("Not a opReturn script.")
	}
	parts = append(parts, Op0Str, OpReturnStr)
	index += 2

	//Chain flag
	if index+2 > limit {
		return nil, false, errors.New("Invalid script data length!")
	}
	icount, lenth := util.DecodeVarIntForScript(scriptBytes[index : index+9])
	if icount == -1 {
		fmt.Println("******************maybe op_Code - ", limit)
		return nil, false, errors.New("Invalid script data length! maybe op_Code")
	}
	index += lenth
	metaLen := icount
	// TODO +len之后 catch fix
	if index+metaLen > limit {
		return nil, false, errors.New("Invalid script data length!")
	}
	isChainFlag, chainFlagStr := IsChainFlag(string(scriptBytes[index : index+metaLen]))
	if !isChainFlag {
		return nil, false, errors.New("Not a metaid script.")
	}
	parts = append(parts, chainFlagStr)
	index += metaLen

	//address
	if index+2 > limit {
		return nil, true, errors.New("Invalid metaid script data length!")
	}
	icount, lenth = util.DecodeVarIntForScript(scriptBytes[index : index+9])
	if icount == -1 {
		fmt.Println("******************maybe op_Code - ", limit)
		return nil, false, errors.New("Invalid script data length! maybe op_Code")
	}
	index += lenth
	addressLen := icount
	address := string(scriptBytes[index : index+addressLen])
	parts = append(parts, address)
	index += addressLen

	//parentTxId
	if index+2 > limit {
		return nil, true, errors.New("Invalid metaid script data length!")
	}
	icount, lenth = util.DecodeVarIntForScript(scriptBytes[index : index+9])
	if icount == -1 {
		fmt.Println("******************maybe op_Code - ", limit)
		return nil, false, errors.New("Invalid script data length! maybe op_Code")
	}
	index += lenth
	parentTxIdLen := icount
	parentTxId := string(scriptBytes[index : index+parentTxIdLen])
	parts = append(parts, parentTxId)
	index += parentTxIdLen

	//data content
	if index == limit {
		return parts, true, nil
	}

	completed := true
	//n := 0
	for completed {
		icount, lenth = util.DecodeVarIntForScript(scriptBytes[index : index+9])
		if icount == -1 {
			fmt.Println("******************maybe op_Code - ", limit)
			return nil, false, errors.New("Invalid script data length! maybe op_Code")
		}
		index += lenth
		partLen := icount
		part := ""
		//if n == 2 && encoding == "binary" {
		//	part = hex.EncodeToString(scriptBytes[index : index+partLen])
		//}else {
		//	part = string(scriptBytes[index : index+partLen])
		//}
		part = hex.EncodeToString(scriptBytes[index : index+partLen])
		parts = append(parts, part)
		index += partLen
		if index == limit {
			completed = false
		}
		//n++
	}

	return parts, true, nil
}
