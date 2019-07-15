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

func FindMACD(values []float64) float64 {

}
