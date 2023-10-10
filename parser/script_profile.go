package parser

const (
	OpCodeHash160     = byte(0xA9)
	OpCodeEqual       = byte(0x87)
	OpCodeEqualVerify = byte(0x88)
	OpCodeCheckSig    = byte(0xAC)
	OpCodeDup         = byte(0x76)
	OpCode_0          = byte(0x00)
	OpCode_1          = byte(0x51)
	OpCheckMultiSig   = byte(0xAE)
	OpPushData1       = byte(0x4C)
	OpPushData2       = byte(0x4D)
	OpPushData3       = byte(0x4E)
	//OpReturn          = byte(0x6A)
)

const (
	OpCodeHash160_Str     = "OP_HASH160"
	OpCodeEqual_Str       = "OP_EQUAL"
	OpCodeEqualVerify_Str = "OP_EQUALVERIFY"
	OpCodeCheckSig_Str    = "OP_CHECKSIG"
	OpCodeDup_Str         = "OP_DUP"
	OpCode_0_Str          = "0"
	OpCode_1_Str          = "1"
	OpCheckMultiSig_Str   = "OP_CHECKMULTISIG"
	OpPushData1_Str       = "OP_PUSHDATA1"
	OpPushData2_Str       = "OP_PUSHDATA2"
	OpPushData3_Str       = "OP_PUSHDATA3"
	OpReturn_Str          = "OP_RETURN"
)
