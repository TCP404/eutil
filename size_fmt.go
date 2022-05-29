package eutil

import "strconv"

const (
	B int64 = 1 << (10*iota + 3)
	KB
	MB
	GB
	TB
	PB
)

// sizeFmt 格式化给定的文件大小，如果文件类型为目录，则返回 "-"
//
// 参数：
//		bit int64: 文件大小的比特数
// 返回值：
// 		string: 格式化后的文件大小，单位 KByte、MByte、GByte、TByte、PByte
func SizeFmt(bit int64) string {
	sizeFloat := float64(bit)
	size := "-"
	unit := "b"
	if bit == 0 {
		return size
	}
	switch {
	case bit < B:
		return strconv.FormatInt(bit, 10) + unit
	case bit >= B && bit < KB:
		sizeFloat /= 1 << 3
		unit = "B"
	case bit >= KB && bit < MB:
		sizeFloat /= 1 << 13
		unit = "KB"
	case bit >= MB && bit < GB:
		sizeFloat /= 1 << 23
		unit = "MB"
	case bit >= GB && bit < TB:
		sizeFloat /= 1 << 33
		unit = "GB"
	case bit >= TB && bit < PB:
		sizeFloat /= 1 << 43
		unit = "TB"
	case bit >= PB:
		sizeFloat /= 1 << 53
		unit = "PB"
	}
	return strconv.FormatFloat(sizeFloat, 'f', 1, 64) + unit
}
