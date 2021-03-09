package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/speps/go-hashids"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %v salt data", os.Args[0])
	}
	data, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("%v", err)
	}

	hd := hashids.NewData()
	hd.Salt = os.Args[1]
	hd.MinLength = 12
	//hd.Alphabet = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRTUVWXYZ2346789"
	hd.Alphabet = "ABCDEFGHJKLMNPQRTUVWXYZ2346789"
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{data})
	fmt.Println(e)

	//d, _ := h.DecodeWithError(e)
	//fmt.Printf("%v-%v-%v %v", e[:4], e[4:8], e[8:], d)
}
