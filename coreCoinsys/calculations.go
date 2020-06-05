package coreCoinsys

import (
	"log"
)

func FindSingleSMA(values []float64, period int) float64 {
	if len(values) < period {
		log.Fatalln("values array is smaller than period, corrupted data")
	}
	sum := float64(0)
	for i := 0; i < period; i++ {
		sum = sum + values[i]
	}
	return sum / float64(period)
}

func FindMultiplier(period int) float64 {
	return (2 / (float64(period) + 1))
}

func FindSingleEMA(iteration int, values []float64, multiplier float64, previousEMA float64) float64 {
	return (values[iteration]-previousEMA)*multiplier + previousEMA
}

func FindEMA(values []float64, period int) []float64 {
	var EMA []float64
	multiplier := FindMultiplier(period)
	j := 1

	// grab initial SMA
	SMA := FindSingleSMA(values, period)
	EMA = append(EMA, SMA)

	// find consecutive EMAs
	EMA = append(EMA, FindSingleEMA(period, values, multiplier, SMA))
	var temp float64
	for _, i := range values[period+1:] {
		temp = ((i - EMA[j]) * multiplier) + EMA[j]
		j = j + 1
		EMA = append(EMA, temp)
	}
	return EMA
}

func FindMACD(values []float64) []float64 {
	var TwelveEMA []float64
	var TwentySixEMA []float64
	var MACD []float64

	TwelveEMA = FindEMA(values, 12)
	TwentySixEMA = FindEMA(values, 26)

	offset := len(TwelveEMA) - len(TwentySixEMA)
	for i := 0; i < len(TwentySixEMA); i++ {
		MACD = append(MACD, TwelveEMA[offset+i]-TwentySixEMA[i])
	}
	return MACD
}

func FindSignalLine(MACDValues []float64) []float64 {
	return FindEMA(MACDValues, 9)
}

func FindHistogram(MACDValues []float64, SignalLineValues []float64) []float64 {
	var histogram []float64
	offset := len(MACDValues) - len(SignalLineValues)
	for i := 0; i < len(SignalLineValues); i++ {
		histogram = append(histogram, MACDValues[offset+i]-SignalLineValues[i])
	}
	return histogram
}

// func BuyOrSellGivenPeriod(HistogramValues []float64, period int) result string {
// 	// use a noise reduction slope calculator
// // possibly use interpolation
// 	var res string
// 	return res
// }
