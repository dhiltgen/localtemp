package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/sirupsen/logrus"

	"github.com/dhiltgen/localtemp/location"
	"github.com/dhiltgen/localtemp/temperature"
)

func main() {
	app := cli.NewApp()
	app.Name = "Local Temperatures"
	app.Usage = `

Gather High,Low temperature information for your current location
Set LAT and LON in your environment to override the geoip lookup for a more accurate predction`
	app.Flags = []cli.Flag{
		/*
		cli.BoolFlag {
			Name: "fahrenheit, f",
			Usage: "Display in F instead of C,",
		},
		*/
		cli.BoolFlag {
			Name: "debug, d",
			Usage: "turn on debug output",
		},
		// TODO add some flags to get future data
	}
	app.Action = func(c *cli.Context) error {
		debug := c.GlobalBool("debug")
		if debug {
			logrus.SetLevel(logrus.DebugLevel)
		}
		l, err := location.GetCurrentLocation()
		if err != nil {
			return err
		}
		t, err := temperature.GetTempData(l)
		if err != nil {
			return err
		}
		fmt.Printf("%0.2f,%0.2f\n", t.High, t.Low)
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
