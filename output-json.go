package main

import (
	"encoding/json"
	"os"
	"time"
)

type RebootRecord struct {
	Generation  string    `json:"generation"`
	Region      string    `json:"region"`
	ServerID    string    `json:"server_uuid"`
	ServerName  string    `json:"server_name"`
	WindowStart time.Time `json:"reboot_window_start"`
	WindowEnd   time.Time `json:"reboot_window_end"`
}

type RebootRecords struct {
	Records []RebootRecord `json:"records"`
}

func outputJSON(entries []entry) {
	enc := json.NewEncoder(os.Stdout)

	var rebootRecords RebootRecords

	rebootRecords.Records = make([]RebootRecord, len(entries))

	for i, e := range entries {

		rebootRecord := RebootRecord{
			Generation:  e.GenType,
			Region:      e.Region,
			ServerID:    e.Server.ID,
			ServerName:  e.Server.Name,
			WindowStart: e.WindowStart,
			WindowEnd:   e.WindowEnd,
		}

		rebootRecords.Records[i] = rebootRecord
	}

	enc.Encode(rebootRecords)

}
