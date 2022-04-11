package main

/**
SHA-1哈希算法
填充消息
被填充消息分组
初始化变量
 */

import (
	"crypto/sha1"
	"fmt"
	"io"
	"math"
	"strconv"
)

var h0 string = "01100111010001010010001100000001"
var h1 string = "11101111110011011010101110001001"
var h2 string = "10011000101110101101110011111110"
var h3 string = "00010000001100100101010001110110"
var h4 string = "11000011110100101110000111110000"

type bits_word struct {
	word [16]string
}

func (b *bits_word) set_word(blockstring string)  {
	//fmt.Println(len(blockstring))
	if len(blockstring) != 512{
		for i:=0;i<16;i++{
			b.word[i] = "nil"
		}
	}else {
		for i:=0;i<16;i++{
			b.word[i] = blockstring[i*32 : i*32+32]
		}
	}
	//fmt.Println(b.word)
}

func (b bits_word) get_word(i int64) string {
	return b.word[i]
}

// 转化字符串为二进制字符串
func stringToBin(s string) (binString string) {
	for _, c := range s{
		binString = fmt.Sprintf("%s%b", binString, c)
	}
	return
}


// 预处理
func pre_deal(raw_str string) []bits_word {
	// 预处理
	bin_string := stringToBin(raw_str)  // 二进制化
	fill_string := padding(bin_string)  // 填充长度
	blocks := group(fill_string)  // 消息分组(按512bits分成块，每块按32bits分成字)
	return blocks

}

// 分步处理过程
func deal(blocks []bits_word, raw_str string) string {
	t := sha1.New()
	io.WriteString(t, raw_str)
	return fmt.Sprintf("%x", t.Sum(nil))
}

// 填充消息
func padding(bin_string string) (fill_string string) {
	/**
	确保消息长度是512bit的倍数
	在原始消息M尾部增加1个“1”和k个0,最终得到新的消息长度必须为512bits的倍数
	 */
	fill_string = bin_string + "1"  // 尾部加上1个1
	//fmt.Println(fill_string)
	
	//计算数组长度并转化为二进制
	len_string := stringToBin(strconv.Itoa(len(bin_string)))
	if len(len_string) > 64 {
		print("可能出错！")
		len_string = len_string[0:63]
	}else {
		// 长度字符串小于等于64bits时补0
		middle_len := len(len_string)
		for i:=0;i<(64-middle_len);i++{
			len_string = "0" + len_string  // 往左补零
		}
	}

	for ;; {
		if math.Mod(float64(len(fill_string)+len(len_string)), 512) == 0 {
			fill_string = fill_string + len_string  // 字符串加上长度字符串
			return
		}else {
			fill_string = fill_string + "0"
		}
	}
}

// 被填充消息分组
func group(fill_string string) []bits_word {
	len_temp := len(fill_string)
	blocks := make([]bits_word, len_temp/512)  // 分块
	for i:=0; i<len(blocks); i++{
		blocks[i].set_word(fill_string[i*512: i*512+512])
	}
	//fmt.Println(blocks)
	return blocks
}




func main() {

	var raw_str string
	fmt.Print("请输入待加密字符串：")
	fmt.Scan(&raw_str)

	// 预处理
	blocks := pre_deal(raw_str)
	//fmt.Println(blocks)

	// 分块处理过程(20*4步)
	fmt.Println(deal(blocks, raw_str))



}
