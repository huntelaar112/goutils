package utils

import "math/big"

func Int64ToBytes(number int64) []byte {
	big := new(big.Int)
	big.SetInt64(number)
	return big.Bytes()
}

func Int64ToString(number int64) string {
	big := new(big.Int)
	big.SetInt64(number)
	return big.String()
}
