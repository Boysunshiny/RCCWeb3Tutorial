package main

import (
	"github.com/afeisir/lsjv-nft-market-bakcend/src/api/router"
)

func main() {
	r := router.NewRouter()
	// fmt.Println(r)
	r.Run(":8811")
}
