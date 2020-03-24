package main

import (
	"stockMarketTool/CryptoPair"
	"stockMarketTool/Quote"
)

func main() {
	quote := Quote.GetQuoteSymbol("AAPL")
	/*Quote.GetRegularMarketData(quote)
	Quote.GetQuoteDepth(quote)
	Quote.GetPreMarketData(quote)
	Quote.GetPostMarketData(quote)
	Quote.GetFiftyTwoWeekRange(quote)
	Quote.GetAverages(quote)
	Quote.GetVolumeMetrics(quote)
	Quote.GetQuoteMetaData(quote)*/
	CryptoPair.GetCryptoPair(quote, )
}



