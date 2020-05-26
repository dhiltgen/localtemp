package temperature

import (
	"fmt"
	"net/http"
	"strconv"

	//"github.com/sirupsen/logrus"
	"github.com/beevik/etree"

	"github.com/dhiltgen/localtemp/location"
)
const (
	BASE_URL="https://graphical.weather.gov/xml/sample_products/browser_interface/ndfdXMLclient.php?&product=time-series&mint=mint&maxt=maxt"
)

type DayTemps struct {
	High float64
	Low float64
	// TODO units?
}

// XML Structure:
// <dwml>
//  <data>
//   <parameters>
//    <temperature type="maximum" ...>
//     <value>xx</value>
//     ...
//    <temperature type="minimum" ...>
//     <value>xx</value>


func GetTempData(loc location.Location) (DayTemps, error) {
	t := DayTemps{}
	resp, err := http.Get(fmt.Sprintf("%s&lat=%0.4f&lon=%0.4f", BASE_URL, loc.Lat, loc.Lon))
	if err != nil {
		return t, err
	}
	defer resp.Body.Close()

	doc := etree.NewDocument()
	_, err = doc.ReadFrom(resp.Body)
	if err != nil {
		return t, err
	}

	for _, e := range doc.FindElements("./dwml/data/parameters/temperature[@type='maximum']/value[0]") {
		t.High, err = strconv.ParseFloat(e.Text(), 64)
		if err != nil {
			return t, err
		}
		break
	}
	for _, e := range doc.FindElements("./dwml/data/parameters/temperature[@type='minimum']/value[0]") {
		t.Low, err = strconv.ParseFloat(e.Text(), 64)
		if err != nil {
			return t, err
		}
		break
	}


	return t, nil
}
