package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"strconv"
)
const TH  = 6

func main() {
	single0 := SingleHash("0")
	single1 := SingleHash("1")
	mltH0 := MultiHash(single0)
	mltH1 := MultiHash(single1)
	println(CombineResults(mltH0, mltH1))
}

func ExecutePipeline() {

}

func SingleHash(data string) string {
	//crcH := StringEncoder(data, false)
	//mdH := StringEncoder(data, true)
	crcH := DataSignerCrc32(data)
	mdH := DataSignerCrc32(DataSignerMd5(data))
	//println(crcH, mdH)
	result := crcH + "~" + mdH
	return result
}

func StringEncoder(data string, isMd bool) string {
	mdH := md5.Sum([]byte(data))
	if isMd {
		mdStr := hex.EncodeToString(mdH[:])
		crcH := crc32.ChecksumIEEE([]byte(mdStr))
		return strconv.FormatUint(uint64(crcH), 10)
	}
	crcH := crc32.ChecksumIEEE([]byte(data))
	return strconv.FormatUint(uint64(crcH), 10)
}

func MultiHash(data string) string {
	var temp string
	for i := 0; i < TH; i++{
		crcLocal := DataSignerCrc32(fmt.Sprintf("%v%v", i, data))
		temp += crcLocal
		println(crcLocal)
	}
	return temp
}


func CombineResults(hash1 string, hash2 string) string {
	return fmt.Sprintf("%v_%v", hash1, hash2)
}
