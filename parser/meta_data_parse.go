package parser

import (
	"github.com/MetaID-Labs/metaid-parser/util"
	"strings"
)

const (
	TextPlain   = "text/plain"
	ImageUrl    = "image/url"
	ImageJpg    = "image/jpg"
	ImageGif    = "image/gif"
	AudioWav    = "audio/wav"
	Transaction = "transaction"
	BsvProtocol = "bsvProtocol"
)

type DataPart struct {
	ChainTag            string      `json:"chainTag"`
	TxId                string      `json:"txID"`
	MetanetId           string      `json:"metanetId"`
	NodePublicKey       string      `json:"nodePublicKey"` //NodePublicKey / NodeAddress
	NodeParentTxId      string      `json:"nodeParentTxId"`
	NodeParentChainFlag string      `json:"nodeParentChainFlag"`
	MetaIdTag           string      `json:"metaIdTag"`
	NodeName            string      `json:"nodeName"`
	Data                interface{} `json:"data"`
	DataString          string      `json:"dataString"` //data raw string
	Encrypt             string      `json:"encrypt"`
	Version             string      `json:"version"`
	DataType            string      `json:"dataType"`
	Encoding            string      `json:"encoding"`
	Params              []string    `json:"params"` //extra params
}

func ChainDataToDataPart(parts []string) *DataPart {
	chainTag := ""
	nodePublicKey := ""
	nodeParentChainTag := ""
	nodeParentTxId := ""
	metaIdTag := ""
	nodeName := ""
	encrypt := ""
	data := ""
	dataString := ""
	version := ""
	dataType := ""
	encoding := ""
	params := make([]string, 0)

	limit := len(parts)
	for index := 0; index < limit; index++ {
		switch index {
		case 0: //OP_0
			break
		case 1: //OP_RETURN
			break
		case 2: //ChainTag
			chainTag = parts[index]
			break
		case 3: //NodeAddress  NodePublicKey
			nodePublicKey = parts[index]
			break
		case 4: //NodeParentTxId
			nodeParentTxId = parts[index]
			break
		case 5: //MetaIdTag
			metaIdTag = strings.ToLower(parts[index])
			break
		case 6: //NodeName
			nodeName = parts[index]
			break
		case 7: //Data
			data = parts[index]
			dataString = parts[index]
			break
		case 8: //Encrypt
			encrypt = parts[index]
			break
		case 9: //Version
			version = parts[index]
			break
		case 10: //DataType
			dataType = parts[index]
			break
		case 11: //Encoding
			encoding = parts[index]
			break
		default:
			params = append(params, parts[index])
		}
	}
	nodeParentChainTag = chainTag
	if strings.Contains(nodeParentTxId, ":") {
		strs := strings.Split(nodeParentTxId, ":")
		if len(strs) == 2 {
			nodeParentChainTag = strs[0]
			nodeParentTxId = strs[1]
		}
	}

	// Compatible with old data
	// Todo
	if len(parts) < 12 {
		if len(parts) == 10 || len(parts) == 11 {
			data = parts[8]
			encrypt = parts[7]
		}
	}

	dataPart := &DataPart{
		ChainTag:            chainTag,
		NodePublicKey:       nodePublicKey,
		NodeParentTxId:      nodeParentTxId,
		NodeParentChainFlag: nodeParentChainTag,
		MetaIdTag:           metaIdTag,
		NodeName:            nodeName,
		Data:                data,
		DataString:          dataString,
		Encrypt:             encrypt,
		Version:             version,
		DataType:            dataType,
		Encoding:            encoding,
		Params:              params,
	}

	// todo fix str data
	str := data
	if strings.HasPrefix(str, "{") && strings.HasSuffix(str, "}") {
		mapData := make(map[string]interface{})
		if err := util.JsonToObject2(str, &mapData); err != nil {
			dataPart.Data = data
		} else {
			dataPart.Data = mapData
		}
	} else if strings.HasPrefix(str, "\"{") && strings.HasSuffix(str, "}\"") {
		if strings.Contains(str, "\\") {
			str = strings.Replace(str, "\\", "", -1)
		}
		str = strings.Trim(str, "\"")
		mapData := make(map[string]interface{})
		if err := util.JsonToObject2(str, &mapData); err != nil {
			dataPart.Data = data
		} else {
			dataPart.Data = mapData
		}
	} else {
		dataPart.Data = data
	}
	return dataPart
}
