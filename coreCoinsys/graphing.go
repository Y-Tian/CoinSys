package coreCoinsys

import (
	"net/http"
	"time"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

func (ga *GraphingAxis) drawChart(res http.ResponseWriter, req *http.Request) {
	var temp []time.Time
	temp = convertEpochToDate(ga.TimestampXAxis)
	var zeroLine []float64
	zeroLine = createZeroLine(ga.TimestampXAxis)
	signalLine := chart.TimeSeries{
		Name:    "Signal Line",
		XValues: temp,
		YValues: ga.SignalYAxis,
	}
	histogramLine := chart.TimeSeries{
		Name:    "Histogram",
		XValues: temp,
		YValues: ga.HistogramYAxis,
	}

	ts1 := chart.TimeSeries{
		Name: "MACD Line",
		Style: chart.Style{
			Show:        true,
			StrokeColor: chart.GetDefaultColor(0),
		},
		XValues: temp,
		YValues: ga.MACDYAxis,
	}
	ts2 := &chart.SMASeries{
		Style: chart.Style{
			Show:        true,
			StrokeColor: drawing.ColorRed,
		},
		InnerSeries: signalLine,
	}
	ts3 := &chart.SMASeries{
		Style: chart.Style{
			Show:        true,
			StrokeColor: drawing.ColorFromHex("989099"),
			FillColor:   drawing.ColorFromHex("989099").WithAlpha(64),
		},
		InnerSeries: histogramLine,
	}
	ts4 := chart.TimeSeries{
		Name: "Zero",
		Style: chart.Style{
			Show:        true,
			StrokeColor: drawing.ColorBlack,
		},
		XValues: temp,
		YValues: zeroLine,
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name:      "Timestamp",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Name:      "MACD Indicator Values",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		Series: []chart.Series{
			ts1,
			ts2,
			ts3,
			ts4,
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

func RunGraph(timestamp []int64, macd []float64, signal []float64, histogram []float64) {
	myGraphHandler := &GraphingAxis{TimestampXAxis: timestamp, MACDYAxis: macd, SignalYAxis: signal, HistogramYAxis: histogram}
	http.HandleFunc("/", myGraphHandler.drawChart)
	http.ListenAndServe(":8080", nil)
}

func convertEpochToDate(timestamp []int64) []time.Time {
	var temp []time.Time
	for _, element := range timestamp {
		temp = append(temp, time.Unix(element, 0))
	}
	return temp
}

func createZeroLine(period []int64) []float64 {
	var temp []float64
	for i := 0; i < len(period); i++ {
		temp = append(temp, 0)
	}
	return temp
}
