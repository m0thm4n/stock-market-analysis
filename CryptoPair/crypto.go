package CryptoPair

import (
	"github.com/piquette/finance-go"
	"stockMarketTool/Quote"
	"fmt"
)

// Get Crypto Pair
func GetCryptoPair(quote *finance.Quote, pair finance.CryptoPair) {
	Quote.GetQuote(quote)
	fmt.Println(quote.Symbol, "Algorithm --> ", pair.Algorithm)
}