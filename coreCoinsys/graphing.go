package coreCoinsys

import (
	"net/http"
	"time"

	"github.com/wcharczuk/go-chart"
)

func (ga *GraphingAxis) drawChart(res http.ResponseWriter, req *http.Request) {
	var temp []time.Time
	temp = convertEpochToDate(ga.XAxis)
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style: chart.StyleShow(),
		},
		Series: []chart.Series{
			chart.TimeSeries{
				XValues: temp,
				YValues: ga.YAxis,
			},
		},
	}

	// mainSeries := chart.ContinuousSeries{
	// 	Name:    "A test series",
	// 	XValues: seq.Range(1.0, 100.0),             //generates a []float64 from 1.0 to 100.0 in 1.0 step increments, or 100 elements.
	// 	YValues: seq.RandomValuesWithMax(100, 100), //generates a []float64 randomly from 0 to 100 with 100 elements.
	// }

	// // note we create a SimpleMovingAverage series by assignin the inner series.
	// // we need to use a reference because `.Render()` needs to modify state within the series.
	// smaSeries := &chart.SMASeries{
	// 	InnerSeries: mainSeries,
	// } // we can optionally set the `WindowSize` property which alters how the moving average is calculated.

	// graph := chart.Chart{
	// 	Series: []chart.Series{
	// 		mainSeries,
	// 		smaSeries,
	// 	},
	// }

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

func RunGraph(x []int64, y []float64) {
	myGraphHandler := &GraphingAxis{XAxis: x, YAxis: y}
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
