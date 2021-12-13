package main

import (
  "net/http"
	"lazy_assets/assetsgo/server"
	pr "lazy_assets/assetsgo/protocol"

)

func main() {
  server := &server.Server{}
  twirpHandler := pr.NewAssetSymbolsServer(server)

  http.ListenAndServe(":8080", twirpHandler)
}