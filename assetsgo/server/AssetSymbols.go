package server

import (
    "context"
    pr "lazy_assets/assetsgo/protocol"

    "fmt"
)

// Server implements the Haberdasher service
type Server struct{}

func (s *Server) GetAssetSymbolTickerPrice(ctx context.Context, request *pr.GetAssetSymbolTickerPriceRequest) (*pr.GetAssetSymbolTickerPriceResponse, error) {
    fmt.Println("work check")
    fmt.Println(request)
    newTicker := &pr.Ticker{
        Ticker: "Fisting is",
        Price: 300.0,
        Currency: "bucks",
    }
    response := &pr.GetAssetSymbolTickerPriceResponse{
        Ticker: newTicker,
    }
    return response, nil
}
