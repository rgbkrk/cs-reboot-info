package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func outputCSV(entries []entry) {
	csvFile, err := os.Create("cs-reboot-info-output.csv")
	if err != nil {
		fmt.Println("Error creating csv file:", err)
		return
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	writer.Write([]string{"generation", "region", "server_uuid", "server_name", "reboot_window_start_UTC", "reboot_window_end_UTC", "reboot_window_start_local", "reboot_window_end_local"})
	for _, e := range entries {
		record := []string{e.GenType, e.Region, e.Server.ID, e.Server.Name, e.WindowStart.String(), e.WindowEnd.String(), e.WindowStart.Local().String(), e.WindowEnd.Local().String()}
		err := writer.Write(record)
		if err != nil {
			fmt.Println("Error writing row to CSV: ", err)
			return
		}
	}
}
