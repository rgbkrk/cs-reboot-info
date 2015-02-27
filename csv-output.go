package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func outputCSV(entries []entry) {
	csvFile, err := os.Create("cs-reboot-info-output.csv")
	if err != nil {
		fmt.Println("Error creating csv file:", err)
		return
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	writer.Write([]string{"Generation", "Region", "Server UUID", "Server Name", "Reboot Window (UTC)", "Reboot Window (Local)"})
	for _, e := range entries {
		utcTime := strings.Join([]string{e.WindowStart.String(), e.WindowEnd.String()}, " - ")
		locTime := strings.Join([]string{e.WindowStart.Local().String(), e.WindowEnd.Local().String()}, " - ")
		record := []string{e.GenType, e.Region, e.Server.ID, e.Server.Name, utcTime, locTime}
		err := writer.Write(record)
		if err != nil {
			fmt.Println("Error writing row to CSV: ", err)
			return
		}
	}
}
