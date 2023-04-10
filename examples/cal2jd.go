package main

import (
	"fmt"
	"os"
	"time"

	"github.com/skrushinsky/scaliger/julian"
)

func main() {
	var dateStr string

	if len(os.Args) > 1 {
		dateStr = os.Args[1]
	} else {
		dateStr = time.Now().Format(time.RFC3339)
	}

	jd, error := julian.DateStringToJulian(dateStr)
	if error != nil {
		fmt.Printf("Invalid date: %s\n. Please, use format: y-mm-ddThh:mm:ssZ ", dateStr)
		os.Exit(1)
	} else {
		fmt.Printf("%.8f\n", jd)
		os.Exit(0)
	}
}
