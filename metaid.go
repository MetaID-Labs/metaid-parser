package metaid

import (
	"fmt"
	"github.com/MetaID-Labs/metaid-base-model/model"
	"github.com/MetaID-Labs/metaid-parser/parser"
	decode "github.com/MetaID-Labs/metaid-parser/transaction"
	"github.com/MetaID-Labs/metaid-parser/util"
	"log"
	"reflect"
)

func DecodeTxToDataPart(bytes []byte) (*parser.DataPart, string, []*model.MetaTxOut, []*model.MetaTxIn, bool) {
	rawTx, err := decode.DecodeRawTransaction(bytes)
	if err != nil {
		log.Printf("decode error %s\n", err)
		return nil, "", nil, nil, false
	}
	vouts := DecodeTxVoutsToMetaVouts(rawTx.TxID, rawTx.Vouts)
	vins := decodeTxVinsToMetaVins(rawTx.Vins)
	for _, out := range rawTx.Vouts {
		data, _, err := parser.DecodeMetaOutScript(out.GetLockScript(), "")
		if err != nil {
			continue
		}
		dataPart := parser.PartsToDataPart(data)
		dataPart.TxId = rawTx.TxID
		dataPart.MetanetId = parser.MakeMetanetId(dataPart.NodePublicKey, dataPart.NodeParentTxId)
		return dataPart, out.GetLockScript(), vouts, vins, true
	}
	return nil, "", nil, nil, false
}

func DecodeTxToCommon(bytes []byte) (string, []*model.MetaTxOut, []*model.MetaTxIn, error) {
	rawTx, err := decode.DecodeRawTransaction(bytes)
	if err != nil {
		log.Printf("decode error %s\n", err)
		return "", nil, nil, err
	}
	vouts := DecodeTxVoutsToMetaVouts(rawTx.TxID, rawTx.Vouts)
	vins := decodeTxVinsToMetaVins(rawTx.Vins)
	return rawTx.TxID, vouts, vins, nil
}

func DecodeTxToDataPartForMetaFile(bytes []byte) (*parser.DataPart, string, bool) {
	rawTx, err := decode.DecodeRawTransaction(bytes)
	if err != nil {
		log.Printf("decode error %s\n", err)
		return nil, "", false
	}
	for _, out := range rawTx.Vouts {
		data, _, err := parser.DecodeMetaOutScript(out.GetLockScript(), "")
		if err != nil {
			continue
		}
		dataPart := parser.PartsToDataPart(data)
		dataPart.TxId = rawTx.TxID
		dataPart.MetanetId = parser.MakeMetanetId(dataPart.NodePublicKey, dataPart.NodeParentTxId)
		return dataPart, out.GetLockScript(), true
	}
	return nil, "", false
}

func DecodeTxVoutsToMetaVouts(txId string, Vouts []decode.TxOut) []*model.MetaTxOut {
	metaVouts := make([]*model.MetaTxOut, 0)
	for _, v := range Vouts {
		scryptType, pkh, err := parser.DecodeNormalOutScript(v.GetLockScript())
		if scryptType != parser.TypeP2PKH {
			continue
		}
		if err != nil {
			fmt.Println("DecodeNormalOutScript err:", err)
			continue
		}
		addr := util.PublicKeyHashToAddress(pkh)

		metaVouts = append(metaVouts, &model.MetaTxOut{
			TxId:         txId,
			Index:        v.GetIndex(),
			Address:      addr,
			PublicKey:    "",
			Value:        v.GetAmount(),
			ScriptPubKey: v.GetLockScript(),
			Type:         "",
		})
	}
	return metaVouts
}

func DecodeTxVoutsToSensibleVouts(txId string, Vouts []decode.TxOut) []*model.MetaTxOut {
	metaVouts := make([]*model.MetaTxOut, 0)
	for _, v := range Vouts {
		metaVouts = append(metaVouts, &model.MetaTxOut{
			TxId:         txId,
			Index:        v.GetIndex(),
			Address:      "",
			PublicKey:    "",
			Value:        v.GetAmount(),
			ScriptPubKey: v.GetLockScript(),
			Type:         "",
		})
	}
	return metaVouts
}

func decodeTxVinsToMetaVins(Vins []decode.TxIn) []*model.MetaTxIn {
	metaVins := make([]*model.MetaTxIn, 0)
	for _, v := range Vins {
		metaVins = append(metaVins, &model.MetaTxIn{
			OutTxId: v.GetTxID(),
			Index:   uint64(v.GetVout()),
		})
	}
	return metaVins
}

func DecodeTxToDataPartForDataString(bytes []byte) (string, bool) {
	rawTx, err := decode.DecodeRawTransaction(bytes)
	if err != nil {
		log.Printf("decode error %s\n", err)
		return "", false
	}
	for _, out := range rawTx.Vouts {
		data, _, err := parser.DecodeMetaOutScript(out.GetLockScript(), "")
		if err != nil {
			continue
		}
		dataPart := parser.PartsToDataPart(data)

		if util.ValueOf(dataPart.Data).Kind() != reflect.String {
			return "", true
		}

		return dataPart.Data.(string), true
	}
	return "", false
}

func DecodeTxToDataPartForDataObj(bytes []byte) (interface{}, bool) {
	rawTx, err := decode.DecodeRawTransaction(bytes)
	if err != nil {
		log.Printf("decode error %s\n", err)
		return "", false
	}
	for _, out := range rawTx.Vouts {
		data, _, err := parser.DecodeMetaOutScript(out.GetLockScript(), "")
		if err != nil {
			continue
		}
		dataPart := parser.PartsToDataPart(data)
		return dataPart.Data, true
	}
	return "", false
}

func DecodeTxToDataPartForImageDataString(bytes []byte) (string, bool) {
	rawTx, err := decode.DecodeRawTransaction(bytes)
	if err != nil {
		log.Printf("decode error %s\n", err)
		return "", false
	}
	for _, out := range rawTx.Vouts {
		data, _, err := parser.DecodeMetaOutScriptForMetaFile(out.GetLockScript(), "")
		if err != nil {
			continue
		}
		dataPart := parser.PartsToDataPart(data)

		if util.ValueOf(dataPart.Data).Kind() != reflect.String {
			return "", true
		}
		return dataPart.Data.(string), true
	}
	return "", false
}
