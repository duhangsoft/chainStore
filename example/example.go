package main

import (
	"fmt"
	"net/http"
)
import (
	"duhangsoft/chainStore"
)

func main() {
	mux := http.NewServeMux()

	chainStore.NewChainStoreServe(mux)
	err := http.ListenAndServe(":8000", mux)
	fmt.Println(err)
}
