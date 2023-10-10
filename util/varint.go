package util

import "math"

// script parse
// 0-10
func DecodeVarIntForScript(buf []byte) (int, int) {
	if buf[0] < 0x4c { //N/A 1-75
		return int(buf[0]), 1
	} else if buf[0] == 0x4c { //OP_PUSHDATA1 - 76
		return int(buf[1]), 2
	} else if buf[0] == 0x4d { //op_pushdata2 - 77
		count :=
			(int(buf[2]) * int(math.Pow(256, 1))) +
				int(buf[1])
		return count, 3
	} else if buf[0] == 0x4e { //OP_PUSHDATA4 - 78
		count :=
			int(buf[4])*int(math.Pow(256, 3)) +
				int(buf[3])*int(math.Pow(256, 2)) +
				int(buf[2])*int(math.Pow(256, 1)) +
				int(buf[1])
		return count, 5
	}
	return -1, -1
}

// tx parse
// 0-9
func DecodeVarIntForTx(buf []byte) (int, int) {
	//if len(buf) != 9 {
	//	return 0, 0
	//}
	if buf[0] <= 0xfc { //252 uint8_t
		return int(buf[0]), 1
	} else if buf[0] == 0xfd { //253 0xFD followed by the length as uint16_t
		return (int(buf[2]) * int(math.Pow(256, 1))) + int(buf[1]), 3
	} else if buf[0] == 0xfe { //254 0xFE followed by the length as uint32_t
		count := (int(buf[4]) * int(math.Pow(256, 3))) +
			(int(buf[3]) * int(math.Pow(256, 2))) +
			(int(buf[2]) * int(math.Pow(256, 1))) +
			int(buf[1])
		return count, 5
	} else if buf[0] == 0xff { //255 0xFF followed by the length as uint64_t
		count := (int(buf[8]) * int(math.Pow(256, 7))) +
			int(buf[7])*int(math.Pow(256, 6)) +
			int(buf[6])*int(math.Pow(256, 5)) +
			int(buf[5])*int(math.Pow(256, 4)) +
			int(buf[4])*int(math.Pow(256, 3)) +
			int(buf[3])*int(math.Pow(256, 2)) +
			int(buf[2])*int(math.Pow(256, 1)) +
			//int(buf[1])*int(math.Pow(256, 1))
			int(buf[1])
		return count, 9
	}
	return 0, 0
}
