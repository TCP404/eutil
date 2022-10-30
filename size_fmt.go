package eutil

import "strconv"

const (
	B int64 = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
)

// sizeFmt 格式化给定的文件大小，如果文件类型为目录，则返回 "-"
//
// 参数：
//
//	bit int64: 文件大小的比特数
//
// 返回值：
//
//	string: 格式化后的文件大小，单位 KByte、MByte、GByte、TByte、PByte
func SizeFmt(bSize int64) string {
	sizeFloat := float64(bSize)
	size := "-"
	unit := "B"

	switch {
	case bSize < B:
		return size
	case bSize >= B && bSize < KB:
		return strconv.FormatInt(bSize, 10) + unit
	case bSize >= KB && bSize < MB:
		sizeFloat /= 1 << 10
		unit = "KB"
	case bSize >= MB && bSize < GB:
		sizeFloat /= 1 << 20
		unit = "MB"
	case bSize >= GB && bSize < TB:
		sizeFloat /= 1 << 30
		unit = "GB"
	case bSize >= TB && bSize < PB:
		sizeFloat /= 1 << 40
		unit = "TB"
	case bSize >= PB:
		sizeFloat /= 1 << 50
		unit = "PB"
	}
	return strconv.FormatFloat(sizeFloat, 'f', 1, 64) + unit
}
