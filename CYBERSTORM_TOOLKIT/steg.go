package main

import (
	"flag"
	"fmt"
	"os"
)

var SENTINEL = []byte{0x00, 0xFF, 0x00, 0x00, 0xFF, 0x00}

func findInterval(wrapper *os.File, hidden *os.File, offset int, sentinel []byte, bitMode bool) int {
	wrapperStat, err := wrapper.Stat()
	if err != nil {
		panic(err)
	}
	var hiddenSize int64
	if hidden != nil {
		hiddenStat, err := hidden.Stat()
		if err != nil {
			panic(err)
		}
		hiddenSize = hiddenStat.Size()
	} else {
		//if there is no hidden file, file size is 0
		hiddenSize = int64(0)
	}

	wrapperSize := wrapperStat.Size()

	if wrapperSize < hiddenSize {
		panic("wrapper file is too small")
	}

	interval := (int(wrapperSize) - offset) / (int(hiddenSize) + len(sentinel))
	if bitMode {
		interval = (int(wrapperSize) - offset) / ((int(hiddenSize) + len(sentinel)) * 8)
	}

	if interval < 1 {
		panic("wrapper file is too small to hide the hidden file in this mode")
	}
	return interval
}

// checks buffer to determine whether or not it matches the sentinal
func matchesSentinel(buf []byte, sentinel []byte) bool {
	if len(buf) < len(sentinel) {
		return false
	}
	for i, v := range sentinel {
		if buf[len(buf)-len(sentinel)+i] != v {
			return false
		}
	}
	return true
}

func byteEncode(file1 []byte, file2 []byte, sentinel []byte, offset int, interval int) {
	for i := 0; i < len(file2); i++ {
		if offset >= len(file1) {
			panic("offset exceeds wrapper file size")
		}
		file1[offset] = file2[i]
		offset += interval
	}

	for i := 0; i < len(sentinel); i++ {
		if offset >= len(file1) {
			panic("offset exceeded wrapper file size while writing sentinel")
		}
		file1[offset] = sentinel[i]
		offset += interval
	}
}

func byteExtract(file []byte, sentinel []byte, offset int, interval int) []byte {
	//create empty byte array for extracted info
	H := make([]byte, 0)

	for i := offset; i < len(file); i += interval {
		H = append(H, file[i])
		//if the current buffer matches any element in the sentinal,  break
		if matchesSentinel(H, sentinel) {
			//remove from byte array
			H = H[:len(H)-len(sentinel)]
			break
		}
	}
	return H
}

func bitEncode(file1 []byte, file2 []byte, sentinel []byte, offset int, interval int) {
	//temp copy of file2 + sentinel to avoid modifying og data
	file2Copy := make([]byte, len(file2))
	copy(file2Copy, file2)

	sentinelCopy := make([]byte, len(sentinel))
	copy(sentinelCopy, sentinel)

	for i := 0; i < len(file2Copy); i++ {
		for j := 0; j < 8; j++ {
			if offset >= len(file1) {
				panic("offset exceeded wrapper file size during bit encoding")
			}
			file1[offset] &= 0xFE
			file1[offset] |= byte((file2Copy[i] >> 7) & 1)
			file2Copy[i] <<= 1
			offset += interval
		}
	}

	for i := 0; i < len(sentinelCopy); i++ {
		for j := 0; j < 8; j++ {
			if offset >= len(file1) {
				panic("offset exceeded wrapper file size while writing sentinel bits")
			}
			file1[offset] &= 0xFE
			file1[offset] |= byte((sentinelCopy[i] >> 7) & 1)
			sentinelCopy[i] <<= 1
			offset += interval
		}
	}
}

func bitExtract(file []byte, sentinel []byte, offset int, interval int) []byte {
	//empty byte array
	H := make([]byte, 0)

	for offset < len(file) {
		var bit byte
		for j := 0; j < 8; j++ {
			if offset >= len(file) {
				break
			}
			bit |= (file[offset] & 1)
			if j < 7 {
				bit <<= 1
				offset += interval
			}
		}
		offset += interval
		//temporarily append bit to H, then compare to sentinal
		H = append(H, bit)
		if matchesSentinel(H, sentinel) {
			//if the current H matches the sentinal, then remove it and continue
			H = H[:len(H)-len(sentinel)]
			break
		}
	}
	return H
}

func main() {
	//declaring flags
	storeFlag := flag.Bool("s", false, "store mode")
	retrieveFlag := flag.Bool("r", false, "retrieve mode")
	bitMode := flag.Bool("b", false, "bit mode")
	byteMode := flag.Bool("B", false, "byte mode")
	offset := flag.Int("o", 0, "offset")
	interval := flag.Int("i", 1, "interval")
	wrapperPath := flag.String("w", "", "wrapper file")
	hiddenPath := flag.String("h", "", "hidden file")

	//parse, validate
	flag.Parse()

	if !*storeFlag && !*retrieveFlag {
		panic("please specify store or retrieve flags")
	}
	if !*bitMode && !*byteMode {
		panic("please specify bit or byte mode")
	}
	if *wrapperPath == "" {
		panic("please specify desired wrapper file")
	}
	if *hiddenPath == "" && *storeFlag {
		panic("please specify message file to hide")
	}

	//open wrapper file
	wrapperFile, err := os.Open(*wrapperPath)
	if err != nil {
		panic(err)
	}
	defer wrapperFile.Close()

	//convert wrapper file into byte array
	wrapperBytes, err := os.ReadFile(*wrapperPath)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stderr, "wrapper bytes: %d\n", len(wrapperBytes))

	if *storeFlag {
		//open hidden file
		hiddenFile, err := os.Open(*hiddenPath)
		if err != nil {
			panic(err)
		}
		defer hiddenFile.Close()

		//convert hidden file into byte array
		hiddenBytes, err := os.ReadFile(*hiddenPath)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(os.Stderr, "hidden bytes: %d\n", len(hiddenBytes))

		//if interval is not specifed, calculate it
		if *interval == 1 {
			*interval = findInterval(wrapperFile, hiddenFile, *offset, SENTINEL, *bitMode)
			fmt.Fprintf(os.Stderr, "calculated interval: %d\n", *interval)
		}

		//check if wrapper file is large enough
		minWrapperSize := (len(hiddenBytes) + len(SENTINEL)) * *interval
		if *bitMode {
			minWrapperSize *= 8
		}

		fmt.Fprintf(os.Stderr, "min wrapper size: %d\n", minWrapperSize)

		if minWrapperSize+*offset > len(wrapperBytes) {
			panic("interval too large. hidden file will not fit in wrapper")
		}

		//encode based on specified mode
		if *byteMode {
			byteEncode(wrapperBytes, hiddenBytes, SENTINEL, *offset, *interval)
		} else {
			bitEncode(wrapperBytes, hiddenBytes, SENTINEL, *offset, *interval)
		}

		os.Stdout.Write(wrapperBytes)

	} else {

		if *interval == 1 {
			fmt.Fprintf(os.Stderr, "no interval specified, using default interval of 2\n")
			*interval = 2 //arbitrary default i chose
		}

		var extracted []byte

		//extract based on specified mode
		if *byteMode {
			extracted = byteExtract(wrapperBytes, SENTINEL, *offset, *interval)
		} else {
			extracted = bitExtract(wrapperBytes, SENTINEL, *offset, *interval)
		}

		os.Stdout.Write(extracted)
	}
}
