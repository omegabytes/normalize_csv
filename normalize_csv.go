package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	Timestamp = iota
	Address
	ZIP
	FullName
	FooDuration
	BarDuration
	TotalDuration
	Notes
)

func main() {
	var out []string = nil

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	header := scanner.Text()

	if !utf8.Valid([]byte(header)) {
		log.Fatal("Malformed string: please double check that the header contains only valid UTF-8 characters")
	}

	out = append(out, header)

	for scanner.Scan() {
		r := csv.NewReader(strings.NewReader(scanner.Text()))
		count := 1
		invalidString := false

		for {
			count++
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			for i := range record {
				if i != Notes {
					if !utf8.Valid([]byte(record[i])) {
						fmt.Fprintf(os.Stderr, "Warning: Malformed string, dropping row %d: %v\n", count, record[i])
						invalidString = true
					}
				}
			}

			if invalidString {
				continue
			}

			record[FullName] = strings.ToUpper(record[FullName])
			record[Timestamp] = convertTimestamp(record[Timestamp])
			record[ZIP] = validateZip(record[ZIP])
			record[FooDuration] = convertDuration(record[FooDuration])
			record[BarDuration] = convertDuration(record[BarDuration])
			record[TotalDuration] = calculateDuration(record[FooDuration], record[BarDuration])

			out = append(out, strings.Join(record, ","))
		}

	}

	if scanner.Err() != nil {
		// handle error.
		log.Fatal(scanner.Err())
	}

	fmt.Print(strings.Join(out, "\n"))

}

// validateZip checks to see if a string is a valid representation of a zip code.
func validateZip(zip string) string {
	if len(zip) < 5 {
		return "0"
	} else if len(zip) > 5 {
		return zip[0:5]
	}
	return zip
}

// convertTimestamp converts a custom US/Pacific timestamp into an ISO-8601 formatted, US/Eastern timestamp
func convertTimestamp(timestamp string) string {
	const InFormat = "1/2/06 3:04:05 PM"
	const OutFormat = "2006-01-02T15:04:05"

	tempTime, err := time.Parse(InFormat, timestamp)
	if err != nil {
		log.Fatal("Unable to parse time")
	}

	tempTime = tempTime.Add(3 * time.Hour)
	convertedTime := tempTime.Format(OutFormat)

	return convertedTime
}

// convertDuration converts a time in HH:MM:SS.MS format to floating-point second format SSSS.MS
func convertDuration(duration string) string {
	dur := strings.Split(duration, ":")

	hours, err := time.ParseDuration(dur[0] + "h")
	mins, err := time.ParseDuration(dur[1] + "m")
	secs, err := time.ParseDuration(dur[2] + "s")

	if err != nil {
		log.Fatal("Unable to convert duration")
	}

	newDur := hours + mins + secs

	return fmt.Sprintf("%g", newDur.Seconds())
}

// calculateDuration finds the amount of time that has passed between a start and end time
func calculateDuration(start string, end string) string {
	s, err := time.ParseDuration(start + "s")
	e, err := time.ParseDuration(end + "s")

	if err != nil {
		log.Fatal("Unable to calculate duration")
	}

	return fmt.Sprintf("%g", (s + e).Seconds())
}
