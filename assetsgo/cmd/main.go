package main

import (
  "fmt"
  "net/http"
	"lazy_assets/assetsgo/server"
	pr "lazy_assets/assetsgo/protocol"

)

func main() {
  server := &server.Server{}
  twirpHandler := pr.NewAssetSymbolsServer(server)

  fmt.Println(twirpHandler)

  http.ListenAndServe(":8080", nil)
}