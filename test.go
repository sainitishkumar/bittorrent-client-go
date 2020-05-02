package main

import (
	"fmt"
	"os"

	"github.com/zeebo/bencode"
)

func main() {
	fmt.Println("Welcome")
	f, _ := os.Open("test.torrent")
	// var reader io.Reader
	fileInfo, _ := f.Stat()
	fileSize := fileInfo.Size()
	bytesvar := make([]byte, fileSize)
	f.Read(bytesvar)
	fmt.Println(string(bytesvar))

	var torrent interface{}
	_ = bencode.DecodeBytes(bytesvar, &torrent)
	torrentData, _ := torrent.(map[string]interface{})
	fmt.Println(torrentData)
}
