package short

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

var chars = [62]string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l",
	"m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x",
	"y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
	"M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X",
	"Y", "Z"}

func GetMd5Code(key string) string {
	m := md5.New()
	m.Write([]byte(key))
	// 加盐
	m.Write([]byte("fengzijk"))

	return hex.EncodeToString(m.Sum(nil))
}

// Get16MD5Encode 返回一个16位md5加密后的字符串
func Get16MD5Encode(data string) string {
	return GetMd5Code(data)[8:24]
}

func GetShortParam(key string) string {

	var res [4]string

	for i := 0; i < len(res); i++ {

		var bt bytes.Buffer
		// 把加密字符按照 8 位一组 16 进制与 0x3FFFFFFF 进行位与运算
		sTempSubString := GetMd5Code(key)[(i * 8):(i*8 + 8)]

		// 这里需要使用 long 型来转换，因为 Inteper .parseInt() 只能处理 31 位 , 首位为符号位 , 如果不用long ，则会越界
		toString, _ := strconv.ParseUint(sTempSubString, 16, 32)

		lHexLong := 0x3FFFFFFF & int64(toString)
		for j := 0; j < 6; j++ {
			// 把得到的值与 0x0000003D 进行位与运算，取得字符数组 chars 索引
			index := 0x0000003D & lHexLong
			// 把取得的字符相加
			bt.WriteString(chars[index])
			// 每次循环按位右移 5 位
			lHexLong = lHexLong >> 5
		}
		// 把字符串存入对应索引的输出数组
		res[i] = bt.String()
	}

	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(3)
	return res[r]
}
