package location


import (
	"os"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	IP_LOOKUP_URL = "https://ipinfo.io/ip"
)


type Location struct {
	Lat float64
	Lon float64
}

func freegeoip(ip string) (Location, error) {
	type locPayload struct {
		Latitude float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
	loc := Location{}
	resp, err := http.Get("https://freegeoip.app/json/"+ip)
	if err != nil {
		return loc, err
	}
	locData, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return loc, err
	}
	//logrus.Debugf("location json: %s", locData)
	l := locPayload{}
	err = json.Unmarshal(locData, &l)
	if err != nil {
		return loc, err
	}
	loc.Lat = l.Latitude
	loc.Lon = l.Longitude
	logrus.Debugf("using geolocation coordinates: %0.2f, %0.2f", loc.Lat, loc.Lon)
	return loc, nil
}


func GetCurrentLocation() (Location, error) {
	lat:= os.Getenv("LAT")
	lon:= os.Getenv("LON")
	loc := Location{}

	if lat!= ""&& lon!= "" {
		data, err := strconv.ParseFloat(lat, 64)
		if err != nil {
			return loc, err
		}
		loc.Lat = data
		data, err = strconv.ParseFloat(lon, 64)
		if err != nil {
			return loc, err
		}
		loc.Lon = data
		logrus.Debugf("using user-specified coordinates: %0.2f, %0.2f", loc.Lat, loc.Lon)
		return loc, nil
	}
	resp, err := http.Get(IP_LOOKUP_URL)
	if err != nil {
		return loc, err
	}
	myip, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return loc, err
	}
	logrus.Debugf("external IP: %s", myip)
	return freegeoip(strings.TrimSpace(string(myip)))
}
