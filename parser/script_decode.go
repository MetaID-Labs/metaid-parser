package parser

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/MetaID-Labs/metaid-parser/util"
)

type ScriptType int64

const (
	TypeEmpty = iota
	TypeP2PK
	TypeP2PKH
	TypeOpreturn
	TypeNotStandard
)

// type p2pk|p2pkh  describe parts error
func DecodeNormalOutScript(scriptPubKey string) (ScriptType, string, error) {
	scriptBytes, err := hex.DecodeString(scriptPubKey)
	if err != nil {
		return TypeEmpty, "", errors.New("Invalid scriptPubKey data!")
	}

	limit := len(scriptBytes)
	if limit == 0 {
		return TypeEmpty, "", errors.New("Invalid script data!")
	}
	index := 0
	icount := 0
	lenth := 0

	//P2PKH
	if len(scriptBytes) == 25 && scriptBytes[0] == OpCodeDup && scriptBytes[1] == OpCodeHash160 && scriptBytes[23] == OpCodeEqualVerify && scriptBytes[24] == OpCodeCheckSig {
		index += 2
		icount, lenth = util.DecodeVarIntForScript(scriptBytes[index : index+9])
		if icount == -1 {
			fmt.Println("******************maybe op_Code - ", limit)
			return TypeEmpty, "", errors.New("Invalid script data length! maybe op_Code")
		}
		index += lenth
		pkhLen := icount
		pkh := hex.EncodeToString(scriptBytes[index : index+pkhLen])
		return TypeP2PKH, pkh, nil
	}

	//P2PK
	if len(scriptBytes) >= 37 && scriptBytes[len(scriptBytes)-1] == OpCodeCheckSig {
		icount, lenth = util.DecodeVarIntForScript(scriptBytes[index : index+9])
		if icount == -1 {
			fmt.Println("******************maybe op_Code - ", limit)
			return TypeEmpty, "", errors.New("Invalid script data length! maybe op_Code")
		}
		index += lenth
		pkLen := icount
		pk := hex.EncodeToString(scriptBytes[index : index+pkLen])
		return TypeP2PK, pk, nil
	}

	//OP_RETURN
	if len(scriptBytes) >= 2 && scriptBytes[0] == OpCode_0 && scriptBytes[1] == OpReturn {
		return TypeOpreturn, "", errors.New("This script data is op_Code")
	}

	//NOT_STANDARD
	return TypeNotStandard, "", nil
}

// type p2pk|p2pkh  describe parts error
func DecodeCommonOutScript(scriptPubKey string) (ScriptType, string, string, []string, error) {
	parts := make([]string, 0)
	opStrs := ""
	scriptBytes, err := hex.DecodeString(scriptPubKey)
	if err != nil {
		return TypeEmpty, "", "", nil, errors.New("Invalid scriptPubKey data!")
	}

	limit := len(scriptBytes)
	if limit == 0 {
		return TypeEmpty, "", "", nil, errors.New("Invalid script data!")
	}
	index := 0

	icount := 0
	lenth := 0

	//P2PKH
	if len(scriptBytes) == 25 && scriptBytes[0] == OpCodeDup && scriptBytes[1] == OpCodeHash160 && scriptBytes[23] == OpCodeEqualVerify && scriptBytes[24] == OpCodeCheckSig {
		index += 2
		icount, lenth = util.DecodeVarIntForScript(scriptBytes[index : index+9])
		if icount == -1 {
			fmt.Println("******************maybe op_Code - ", limit)
			return TypeEmpty, "", "", nil, errors.New("Invalid script data length! maybe op_Code")
		}
		index += lenth
		pkhLen := icount
		pkh := hex.EncodeToString(scriptBytes[index : index+pkhLen])
		parts = append(parts, OpCodeDup_Str, OpCodeHash160_Str, pkh, OpCodeEqualVerify_Str)
		opStrs = partsToString(parts)
		return TypeP2PKH, pkh, opStrs, parts, nil
	}

	//P2PK
	if len(scriptBytes) >= 37 && scriptBytes[len(scriptBytes)-1] == OpCodeCheckSig {
		icount, lenth = util.DecodeVarIntForScript(scriptBytes[index : index+9])
		if icount == -1 {
			fmt.Println("******************maybe op_Code - ", limit)
			return TypeEmpty, "", "", nil, errors.New("Invalid script data length! maybe op_Code")
		}
		index += lenth
		pkLen := icount
		pk := hex.EncodeToString(scriptBytes[index : index+pkLen])
		parts = append(parts, pk, OpCodeEqualVerify_Str)
		opStrs = partsToString(parts)
		return TypeP2PK, pk, opStrs, parts, nil
	}

	//OP_RETURN
	if scriptBytes[0] == OpCode_0 && scriptBytes[1] == OpReturn {
		index += 2
		completed := true
		for completed {
			icount, lenth = util.DecodeVarIntForScript(scriptBytes[index : index+9])
			if icount == -1 {
				fmt.Println("******************maybe op_Code - ", limit)
				return TypeEmpty, "", "", nil, errors.New("Invalid script data length! maybe op_Code")
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
		opStrs = partsToString(parts)
		return TypeOpreturn, "", opStrs, parts, nil
	}

	//NOT_STANDARD
	return TypeNotStandard, "", "", nil, nil
}

func partsToString(parts []string) string {
	opStrs := ""
	for k, v := range parts {
		if k == 0 {
			opStrs = v
		} else {
			opStrs = opStrs + " " + v
		}
	}
	return opStrs
}
