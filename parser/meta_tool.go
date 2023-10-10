package parser

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var _BeginMetaBlockHeight = int64(680000)
var _PerMetaBlockCount = int64(144)

func init() {

}

/*
*

	Make MetanetID for finding unique parent node
*/
func MakeMetanetId(nodePubKey string, nodeParentTxId string) string {
	if nodeParentTxId == "" || nodePubKey == "" {
		return ""
	}
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%s%s", nodePubKey, nodeParentTxId)))
	bytes := hash.Sum(nil)
	hashCode := hex.EncodeToString(bytes)
	return hashCode
}

/*
*

	Make NodeID
*/
func MakeNodeId(nodePubKey string, nodeTxId string) string {
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%s%s", nodePubKey, nodeTxId)))
	bytes := hash.Sum(nil)
	hashCode := hex.EncodeToString(bytes)
	return hashCode
}

func MakeMetaBlockHeight(blockHeight int64) int64 {
	if blockHeight <= 0 {
		return -1
	}
	return ((blockHeight - _BeginMetaBlockHeight) / _PerMetaBlockCount) + 1
}

func GetStartBlockHeightAndEndBlockHeight(metaBlockHeight int64) (int64, int64) {
	var (
		startBlockHeight int64 = 0
		endBlockHeight   int64 = 0
	)
	startBlockHeight = GetMetaBlockStartBlockHeight(metaBlockHeight)
	endBlockHeight = GetMetaBlockEndBlockHeight(metaBlockHeight)
	return startBlockHeight, endBlockHeight
}

func GetMetaBlockStartBlockHeight(metaBlockHeight int64) int64 {
	//if !(conf.IsTestMVC() || conf.IsTestPreMVC()) {
	//	if metaBlockHeight <= 0 {
	//		return -1
	//	}
	//}
	return ((metaBlockHeight - 1) * _PerMetaBlockCount) + _BeginMetaBlockHeight
}

func GetMetaBlockEndBlockHeight(metaBlockHeight int64) int64 {
	//if !(conf.IsTestMVC() || conf.IsTestPreMVC()) {
	//	if metaBlockHeight <= 0 {
	//		return -1
	//	}
	//}
	return ((metaBlockHeight) * _PerMetaBlockCount) + _BeginMetaBlockHeight
}
