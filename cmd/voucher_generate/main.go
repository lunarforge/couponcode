package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/lunarforge/couponcode/pkg/voucher"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %v seed", os.Args[0])
	}
	seedInt, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatalf("%v", err)
	}

	rand.Seed(seedInt)
	//cc := voucher.New(3, 5)
	code := voucher.Generate()
	validated, err := voucher.Validate(code)
	if err != nil {
		log.Fatalf("%v %v %v\n", code, validated, err)
	}
	fmt.Printf("%v", code)

}
