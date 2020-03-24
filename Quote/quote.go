package Quote

import (
	"fmt"
	"github.com/piquette/finance-go"
	quote2 "github.com/piquette/finance-go/quote"
)

// Get Quote type
func GetQuoteSymbol(symbol string) *finance.Quote {
	quote, err := quote2.Get(symbol)
	if err != nil {
		fmt.Println(err)
	}

	return quote
}

// Get Regular session quote data
func GetRegularMarketData(quote *finance.Quote) {
	fmt.Println(quote.Symbol, "Regular Market Change Percent -->", quote.RegularMarketChangePercent)
	fmt.Println(quote.Symbol, "Regular Market Previous Close -->", quote.RegularMarketPreviousClose)
	fmt.Println(quote.Symbol, "Regular Market Previous Close -->", quote.RegularMarketPreviousClose)
	fmt.Println(quote.Symbol, "Regular Market Previous Close -->", quote.RegularMarketPreviousClose)
	fmt.Println(quote.Symbol, "Regular Market Price -->", quote.RegularMarketPrice)
	fmt.Println(quote.Symbol, "Regular Market Time -->", quote.RegularMarketTime)
	fmt.Println(quote.Symbol, "Regular Market Change -->", quote.RegularMarketChange)
	fmt.Println(quote.Symbol, "Regular Market Open -->", quote.RegularMarketOpen)
	fmt.Println(quote.Symbol, "Regular Market Day High -->", quote.RegularMarketDayHigh)
	fmt.Println(quote.Symbol, "Regular Market Day Low -->", quote.RegularMarketDayLow)
	fmt.Println(quote.Symbol, "Regular Market Volume -->", quote.RegularMarketVolume)
}

// Get Quote Depth
func GetQuoteDepth(quote *finance.Quote) {
	fmt.Println(quote.Symbol, "Bid price --> ", quote.Bid)
	fmt.Println(quote.Symbol, "Ask price --> ", quote.Ask)
	fmt.Println(quote.Symbol, "BidSize --> ", quote.BidSize)
	fmt.Println(quote.Symbol, "AskSize --> ", quote.AskSize)
}

// Get Pre-Market Data
func GetPreMarketData(quote *finance.Quote) {
	fmt.Println(quote.Symbol, "Pre-Market Price --> ", quote.PreMarketPrice)
	fmt.Println(quote.Symbol, "Pre-Market Change --> ", quote.PreMarketChange)
	fmt.Println(quote.Symbol, "Pre-Market Change Percent --> ", quote.PreMarketChangePercent)
	fmt.Println(quote.Symbol, "Pre-Market Time --> ", quote.PreMarketTime)
}

// Get Post-Market Data
func GetPostMarketData(quote *finance.Quote) {
	fmt.Println(quote.Symbol, "Post-Market Price --> ", quote.PostMarketPrice)
	fmt.Println(quote.Symbol, "Post-Market Change --> ", quote.PostMarketChange)
	fmt.Println(quote.Symbol, "Post-Market Change Percent --> ", quote.PostMarketChangePercent)
	fmt.Println(quote.Symbol, "Post-Market Time --> ", quote.PostMarketTime)
}

// Get 52wk ranges
func GetFiftyTwoWeekRange(quote *finance.Quote) {
	fmt.Println(quote.Symbol, "Fifty Two Week Low Change --> ", quote.FiftyTwoWeekLowChange)
	fmt.Println(quote.Symbol, "Fifty Two Week Change Percent --> ", quote.FiftyTwoWeekLowChangePercent)
	fmt.Println(quote.Symbol, "Fifty Two Week High Change --> ", quote.FiftyTwoWeekHighChange)
	fmt.Println(quote.Symbol, "Fifty Two Week High Change Percent --> ", quote.FiftyTwoWeekHighChangePercent)
	fmt.Println(quote.Symbol, "Fifty Two Week High --> ", quote.FiftyTwoWeekHigh)
	fmt.Println(quote.Symbol, "Fifty Two Week Low --> ", quote.FiftyTwoWeekLow)
}

// Get Averages
func GetAverages(quote *finance.Quote) {
	fmt.Println(quote.Symbol, "Fifty Day Average --> ", quote.FiftyDayAverage)
	fmt.Println(quote.Symbol, "Fifty Day Average Change --> ", quote.FiftyDayAverageChange)
	fmt.Println(quote.Symbol, "Fifty Day Average Change Percent --> ", quote.FiftyDayAverageChangePercent)
	fmt.Println(quote.Symbol, "Two Hundred Day Average --> ", quote.TwoHundredDayAverage)
	fmt.Println(quote.Symbol, "Two Hundred Day Average Change --> ", quote.TwoHundredDayAverageChange)
	fmt.Println(quote.Symbol, "Two Hundred Day Average Change Percent --> ", quote.TwoHundredDayAverageChangePercent)
}

func GetVolumeMetrics(quote *finance.Quote) {
	fmt.Println(quote.Symbol, "Average Daily Volume 3 Month --> ", quote.AverageDailyVolume3Month)
	fmt.Println(quote.Symbol, "Average Daily Volume 10 Day --> ", quote.AverageDailyVolume10Day)
}

func GetQuoteMetaData(quote *finance.Quote) {
	fmt.Println(quote.Symbol, "Quote Source --> ", quote.QuoteSource)
	fmt.Println(quote.Symbol, "Quote Currency ID --> ", quote.CurrencyID)
	fmt.Println(quote.Symbol, "Is Tradeable --> ", quote.IsTradeable)
	fmt.Println(quote.Symbol, "Quote Delay --> ", quote.QuoteDelay)
	fmt.Println(quote.Symbol, "Full Exchange Name --> ", quote.FullExchangeName)
	fmt.Println(quote.Symbol, "Source Interval", quote.SourceInterval)
	fmt.Println(quote.Symbol, "Exchange Timezone Name --> ", quote.ExchangeTimezoneName)
	fmt.Println(quote.Symbol, "Exchange Timezone Short Name --> ", quote.ExchangeTimezoneShortName)
	fmt.Println(quote.Symbol, "GMT Off Set Milliseconds --> ", quote.GMTOffSetMilliseconds)
	fmt.Println(quote.Symbol, "Market ID --> ", quote.MarketID)
	fmt.Println(quote.Symbol, "Exchange ID --> ", quote.ExchangeID)
}

func GetQuote(quote *finance.Quote) {
	GetRegularMarketData(quote)
	GetQuoteDepth(quote)
	GetPreMarketData(quote)
	GetPostMarketData(quote)
	GetFiftyTwoWeekRange(quote)
	GetAverages(quote)
	GetVolumeMetrics(quote)
	GetQuoteMetaData(quote)
}