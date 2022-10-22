package main

import (
	"fmt"
)

const TH = 6

func main() {
	//single0 := SingleHash("0")
	//single1 := SingleHash("1")
	//mltH0 := MultiHash(single0)
	//mltH1 := MultiHash(single1)
	//println(CombineResults(mltH0, mltH1))
}

func ExecutePipeline() {

}

func SingleHash(in, out chan interface{}) {
	//data :=  <- in
	data := "0"
	crcH := DataSignerCrc32(data)
	mdH := DataSignerCrc32(DataSignerMd5(data))
	println(fmt.Sprintf("%v~%v", crcH, mdH))
}

func MultiHash(in, out chan interface{}) string {
	data := "0"
	var temp string
	for i := 0; i < TH; i++ {
		crcLocal := DataSignerCrc32(fmt.Sprintf("%v%v", i, data))
		temp += crcLocal
		println(crcLocal)
	}
	return temp
}

func CombineResults(in, out chan interface{}) string {
	hash1 := "dfsf"
	hash2 := "dfkjsd"
	return fmt.Sprintf("%v_%v", hash1, hash2)
}


//func SingleHash(data string) string {
//	//crcH := StringEncoder(data, false)
//	//mdH := StringEncoder(data, true)
//	crcH := DataSignerCrc32(data)
//	mdH := DataSignerCrc32(DataSignerMd5(data))
//	//println(crcH, mdH)
//	return fmt.Sprintf("%v~%v", crcH, mdH)
//}
