package bytesex

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"io/ioutil"
)
func BytesToInt32(b []byte, start int) int {
	return int(b[start]<<24 + b[start+1]<<16 + b[start+2]<<8 + b[start+3])
}
func HexStringToBytes(s string) []byte {
	n := len(s)
	byteLength := n / 2

	r := make([]byte, byteLength)

	i := 0
	j := 0
	b := (byte)(0)
	for (i < n) && (j < byteLength) {
		switch s[i] {
		case '1':
			b = (byte)(1)
		case '2':
			b = (byte)(2)
		case '3':
			b = (byte)(3)
		case '4':
			b = (byte)(4)
		case '5':
			b = (byte)(5)
		case '6':
			b = (byte)(6)
		case '7':
			b = (byte)(7)
		case '8':
			b = (byte)(8)
		case '9':
			b = (byte)(9)
		case 'a', 'A':
			b = (byte)(10)
		case 'b', 'B':
			b = (byte)(11)
		case 'c', 'C':
			b = (byte)(12)
		case 'd', 'D':
			b = (byte)(13)
		case 'e', 'E':
			b = (byte)(14)
		case 'f', 'F':
			b = (byte)(15)
		default:
			b = (byte)(0)
		}
		i++

		if i%2 == 1 {
			r[j] = b
		} else {
			r[j] = r[j] << 4
			r[j] = r[j] | b
			j++
		}
	}

	return r
}
func ZlibZipBytes(input []byte) ([]byte, error) {
	var buf bytes.Buffer
	compressor, err := zlib.NewWriterLevel(&buf, zlib.BestCompression)
	if err != nil {
		return nil, err
	}
	compressor.Write(input)
	compressor.Close()
	return buf.Bytes(), nil
}

func ZlibUnzipBytes(input []byte) ([]byte, error) {
	b := bytes.NewReader(input)
	r, err := zlib.NewReader(b)
	defer r.Close()
	if err != nil {
		return nil, err
	}
	data, _ := ioutil.ReadAll(r)
	return data, nil
}

// 混淆[]byte
func ConfusedTwo(sourceBytes []byte) []byte {
	var confusedBytes []byte = make([]byte, len(sourceBytes))
	idx := 0
	for index := 0; index < len(sourceBytes); index++ {
		if index%2 == 0 {
			confusedBytes[idx] = byte(255 - sourceBytes[index])
			idx++
		}
	}
	for index := 0; index < len(sourceBytes); index++ {
		if index%2 == 1 {
			confusedBytes[idx] = byte(255 - sourceBytes[index])
			idx++
		}
	}
	return confusedBytes
}

// 反混淆[]byte
func UnConfusedTwo(confusedBytes []byte) []byte {
	loopCount := len(confusedBytes)
	count := loopCount / 2
	if loopCount%2 == 1 {
		count = loopCount/2 + 1
	}
	beforeBytes := confusedBytes[0:count]
	afterBytes := confusedBytes[count:]
	var unConfusedBytes []byte = make([]byte, loopCount)
	beforeIndex := 0
	afterIndex := 0
	for index := 0; index < loopCount; index++ {
		if index%2 == 0 {
			if beforeIndex >= len(beforeBytes) {
				continue
			}
			unConfusedBytes[index] = 255 - beforeBytes[beforeIndex]
			beforeIndex++
		} else {
			if afterIndex >= len(afterBytes) {
				continue
			}
			unConfusedBytes[index] = 255 - afterBytes[afterIndex]
			afterIndex++
		}
	}
	return unConfusedBytes
}

//反转[]byte
func ReversalBytes(source []byte) []byte {
	builder := make([]byte, len(source))
	for i, v := range source {
		builder[i] = 255 - v
	}
	return builder
}

//把十六进制字符串转成字节数组
//func HexStringToBytes(hex_str string) []byte {
//	len := len(hex_str) / 2
//	byte_array := make([]byte, len)
//	for index := 0; index < len; index++ {
//		ina, _ := strconv.ParseInt((hex_str, index*2, 2), 16, 32)
//		byte_array[index] = byte(ina)
//	}
//	return byte_array
//}

//整形转换成字节
func UIntToBytes(value uint32) []byte {
	arr := []byte{0, 0, 0, 0}
	arr[3] = byte((value >> 24) & 0xff)
	arr[2] = byte((value >> 16) & 0xff)
	arr[1] = byte((value >> 8) & 0xff)
	arr[0] = byte(value & 0xff)
	return arr
}

//字节转换成整形
func BytesToUInt(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}