package main

import (
	"flag"
	"fmt"
	"github.com/vjeantet/jodaTime"
	"os"
	"time"
)

func main() {
	format := "yyyy-MM-dd"
	var help bool

	flag.StringVar(&format, "f", format, "Please inform date pattern.. More information https://github.com/vjeantet/jodaTime#format")
	flag.BoolVar(&help, "h", false, "Show Usage")
	flag.Parse()

	if help == true {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	fmt.Print(jodaTime.Format(format, time.Now()))
}
