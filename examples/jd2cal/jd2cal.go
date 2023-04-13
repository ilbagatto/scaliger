package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/skrushinsky/scaliger/julian"
)

func main() {
	var jd float64
	var err error

	if len(os.Args) > 1 {
		jdStr := os.Args[1]
		jd, err = strconv.ParseFloat(jdStr, 32)
		if err != nil {
			fmt.Printf("Invalid Julian Date: %s\n.", jdStr)
			os.Exit(1)
		}

	} else {
		fmt.Print("Usage: jd2cal JD\n")
		os.Exit(1)
	}

	dateStr := julian.JulianToDateString(jd)
	fmt.Printf("%s\n", dateStr)
	os.Exit(0)
}
