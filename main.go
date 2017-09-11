// Good times.

package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func main() {
	command := flag.String("command", "", "Script to be repeated.")
	start := flag.String("start", "", "Start date in 2006-01-02 format.")
	end := flag.String("end", "", "End date in 2006-01-02 format.")
	params := flag.String("params", "", "Params to pass to the command with %s being the date we're using: dater --command=\"path/to/shipper\" --params=\"-v -d=%s\"")
	flag.Parse()

	if *command == "" {
		fmt.Println("--command is required")
		return
	}
	if *params == "" {
		fmt.Println("--params is required with a minimum of %s")
		return
	}
	if *start == "" {
		fmt.Println("--start is required in format of 2006-01-02")
		return
	}
	if *end == "" {
		fmt.Println("--end is required in format of 2006-01-02")
		return
	}

	if *start > *end {
		fmt.Println("The start date cannot be after the end.")
		return
	}

	t1, err := time.Parse("2006-01-02", *start)
	if err != nil {
		fmt.Printf("error encountered while parsing the start date: %s", err)
		return
	}

	t2, err := time.Parse("2006-01-02", *end)
	if err != nil {
		fmt.Printf("error encountered while parsing the end date: %s", err)
		return
	}

	parts := strings.Fields(*params)
	parts = parts[1:len(parts)]

	t2 = t2.Add(24 * time.Hour)
	for !t1.Equal(t2) {
		args := strings.Fields(fmt.Sprintf(*params, t1.Format("2006-01-02")))
		cmdOut, err := exec.Command(*command, args...).Output()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(cmdOut))

		t1 = t1.Add(24 * time.Hour)
	}
}
