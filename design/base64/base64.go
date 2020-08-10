package base64

const (
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

var base64TableMap = make(map[rune]int, len(base64Table))

func init() {
	for i, b := range base64Table {
		base64TableMap[b] = i
	}
}

func Encode(src []byte) []byte {

	// 1.补零操作
	fill := -1
	sLen := len(src)
	if sLen%3 != 0 {
		fill = 3 - (sLen % 3)
	}
	fillTemp := fill
	for fill > 0 {
		src = append(src, 0x00)
		fill--
	}

	// 2.转6bit，高位补0
	len6bit := len(src) * 8 / 6
	sixBitGroup := make([]byte, len6bit)

	for i, j := 0, 0; i < len6bit; i, j = i+4, j+3 {
		sixBitGroup[i] = ((src[j] & 0xFC) >> 2)
		sixBitGroup[i+1] = ((src[j] & 0x03) << 4) + ((src[j+1] & 0xF0) >> 4)
		sixBitGroup[i+2] = ((src[j+1] & 0x0F) << 2) + ((src[j+2] & 0xC0) >> 6)
		sixBitGroup[i+3] = (src[j+2] & 0x3F)
	}

	// 3.base64编码
	for i, b := range sixBitGroup {
		sixBitGroup[i] = base64Table[b]
	}

	// 4.将第一步填充的补位0置换为=
	for i := 0; i < fillTemp; i++ {
		sixBitGroup[len6bit-i-1] = '='
	}
	return sixBitGroup
}

func Decode(src []byte) []byte {
	// 补位数量
	var fillNum int
	// 1. 置换为正常数值
	for i, b := range src {
		if b == '=' {
			fillNum++
			src[i] = base64Table[0]
			continue
		}
		src[i] = byte(base64TableMap[rune(b)])
	}

	times := 0
	// 2. 将8bit转化为6bit
	for i, j := 0, 0; j < len(src); i, j = i+3, j+4 {
		src[i] = (src[j] << 2) + ((src[j+1] & 0xF0) >> 4)
		src[i+1] = ((src[j+1] & 0x0F) << 4) + ((src[j+2] & 0xFC) >> 2)
		src[i+2] = ((src[j+2] & 0x03) << 6) + src[j+3]
		times++
	}
	return src[:len(src)-times-fillNum]
}
