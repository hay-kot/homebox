// Package reporting provides a way to import CSV files into the database.
package reporting

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io"
	"strings"
)

var (
	ErrNoHomeboxHeaders       = errors.New("no headers found")
	ErrMissingRequiredHeaders = errors.New("missing required headers `HB.location` or `HB.name`")
)

// determineSeparator determines the separator used in the CSV file
// It returns the separator as a rune and an error if it could not be determined
//
// It is assumed that the first row is the header row and that the separator is the same
// for all rows.
//
// Supported separators are `,` and `\t`
func determineSeparator(data []byte) (rune, error) {
	// First row
	firstRow := bytes.Split(data, []byte("\n"))[0]

	// find first comma or /t
	comma := bytes.IndexByte(firstRow, ',')
	tab := bytes.IndexByte(firstRow, '\t')

	switch {
	case comma == -1 && tab == -1:
		return 0, errors.New("could not determine separator")
	case tab > comma:
		return '\t', nil
	default:
		return ',', nil
	}
}

// readRawCsv reads a CSV file and returns the raw data as a 2D string array
// It determines the separator used in the CSV file and returns an error if
// it could not be determined
func readRawCsv(r io.Reader) ([][]string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(bytes.NewReader(data))

	// Determine separator
	sep, err := determineSeparator(data)
	if err != nil {
		return nil, err
	}

	reader.Comma = sep

	return reader.ReadAll()
}

// parseHeaders parses the homebox headers from the CSV file and returns a map of the headers
// and their column index as well as a list of the field headers (HB.field.*) in the order
// they appear in the CSV file
//
// It returns an error if no homebox headers are found
func parseHeaders(headers []string) (hbHeaders map[string]int, fieldHeaders []string, err error) {
	hbHeaders = map[string]int{} // initialize map

	for col, h := range headers {
		if strings.HasPrefix(h, "HB.field.") {
			fieldHeaders = append(fieldHeaders, h)
		}

		if strings.HasPrefix(h, "HB.") {
			hbHeaders[h] = col
		}
	}

	required := []string{"HB.location", "HB.name"}
	for _, h := range required {
		if _, ok := hbHeaders[h]; !ok {
			return nil, nil, ErrMissingRequiredHeaders
		}
	}

	if len(hbHeaders) == 0 {
		return nil, nil, ErrNoHomeboxHeaders
	}

	return hbHeaders, fieldHeaders, nil
}
