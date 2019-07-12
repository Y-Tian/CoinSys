package coreCoinsys

func FindSingleSMA(values []float64, period int) float64 {
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

func FindTotalEMA(values []float64, period int) float64 {
	multiplier := FindMultiplier(period)
	initialSMA := FindSingleSMA(values, period)
	EMA := FindSingleEMA(0, values, multiplier, initialSMA)
	for i := 1; i < period; i++ {
		EMA = FindSingleEMA(i, values, multiplier, EMA)
	}
	return EMA
}

// need all historical values to run forumla, need to return a macd array
// must pull all data from coin history

func FindMACD(values []float64) float64 {

}
