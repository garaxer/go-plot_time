package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func getValues() []float64 {

	return []float64{2021, 2021, 2020, 2020, 2019, 2019, 2019, 2019, 2019, 2018, 2017, 2017, 2017, 2017, 2017, 2017, 2016, 2016, 2016, 2015, 2014, 2014, 2014, 2014, 2013, 2013, 2012, 2012, 2011, 2011, 2011, 2010, 2010, 2009, 2007, 2006, 2006, 2006, 2005, 2000, 1999, 1997, 1996, 1993, 1993, 1990, 1987, 1984, 1983}
}

func main() {
	//make data
	//var values plotter.Values
	values := getValues()
	histPlot(values)
}

// [...s.matchAll(/	([0-9]{4})	/g)].map(x => parseInt(x[1]))

func histPlot(values plotter.Values) {
	p := plot.New()

	p.Title.Text = "Time is increasing"

	hist, err := plotter.NewHist(values, 25)
	if err != nil {
		panic(err)
	}
	p.Add(hist)

	if err := p.Save(3*vg.Inch, 3*vg.Inch, "hist2.png"); err != nil {
		panic(err)
	}

}
