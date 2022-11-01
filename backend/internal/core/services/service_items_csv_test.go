package services

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const CSV_DATA = `
Import Ref,Location,Labels,Quantity,Name,Description,Insured,Serial Number,Mode Number,Manufacturer,Notes,Purchase From,Purchased Price,Purchased Time,Lifetime Warranty,Warranty Expires,Warranty Details,Sold To,Sold Price,Sold Time,Sold Notes
A,Garage,IOT;Home Assistant; Z-Wave,1,Zooz Universal Relay ZEN17,Description 1,TRUE,,ZEN17,Zooz,,Amazon,39.95,10/13/2021,,10/13/2021,,,,10/13/2021,
B,Living Room,IOT;Home Assistant; Z-Wave,1,Zooz Motion Sensor,Description 2,FALSE,,ZSE18,Zooz,,Amazon,29.95,10/15/2021,,10/15/2021,,,,10/15/2021,
C,Office,IOT;Home Assistant; Z-Wave,1,Zooz 110v Power Switch,Description 3,TRUE,,ZEN15,Zooz,,Amazon,39.95,10/13/2021,,10/13/2021,,,,10/13/2021,
D,Downstairs,IOT;Home Assistant; Z-Wave,1,Ecolink Z-Wave PIR Motion Sensor,Description 4,FALSE,,PIRZWAVE2.5-ECO,Ecolink,,Amazon,35.58,10/21/2020,,10/21/2020,,,,10/21/2020,
E,Entry,IOT;Home Assistant; Z-Wave,1,Yale Security Touchscreen Deadbolt,Description 5,TRUE,,YRD226ZW2619,Yale,,Amazon,120.39,10/14/2020,,10/14/2020,,,,10/14/2020,
F,Kitchen,IOT;Home Assistant; Z-Wave,1,Smart Rocker Light Dimmer,Description 6,FALSE,,39351,Honeywell,,Amazon,65.98,09/30/2020,,09/30/2020,,,,09/30/2020,`

func loadcsv() [][]string {
	reader := csv.NewReader(bytes.NewBuffer([]byte(CSV_DATA)))

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	return records
}

func Test_CorrectDateParsing(t *testing.T) {
	t.Parallel()

	expected := []time.Time{
		time.Date(2021, 10, 13, 0, 0, 0, 0, time.UTC),
		time.Date(2021, 10, 15, 0, 0, 0, 0, time.UTC),
		time.Date(2021, 10, 13, 0, 0, 0, 0, time.UTC),
		time.Date(2020, 10, 21, 0, 0, 0, 0, time.UTC),
		time.Date(2020, 10, 14, 0, 0, 0, 0, time.UTC),
		time.Date(2020, 9, 30, 0, 0, 0, 0, time.UTC),
	}

	records := loadcsv()

	for i, record := range records {
		if i == 0 {
			continue
		}
		entity := newCsvRow(record)
		expected := expected[i-1]

		assert.Equal(t, expected, entity.Item.PurchaseTime, fmt.Sprintf("Failed on row %d", i))
		assert.Equal(t, expected, entity.Item.WarrantyExpires, fmt.Sprintf("Failed on row %d", i))
		assert.Equal(t, expected, entity.Item.SoldTime, fmt.Sprintf("Failed on row %d", i))
	}
}

func Test_csvRow_getLabels(t *testing.T) {
	type fields struct {
		LabelStr string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "basic test",
			fields: fields{
				LabelStr: "IOT;Home Assistant;Z-Wave",
			},
			want: []string{"IOT", "Home Assistant", "Z-Wave"},
		},
		{
			name: "no labels",
			fields: fields{
				LabelStr: "",
			},
			want: []string{},
		},
		{
			name: "single label",
			fields: fields{
				LabelStr: "IOT",
			},
			want: []string{"IOT"},
		},
		{
			name: "trailing semicolon",
			fields: fields{
				LabelStr: "IOT;",
			},
			want: []string{"IOT"},
		},

		{
			name: "whitespace",
			fields: fields{
				LabelStr: " IOT;		Home Assistant;   Z-Wave ",
			},
			want: []string{"IOT", "Home Assistant", "Z-Wave"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := csvRow{
				LabelStr: tt.fields.LabelStr,
			}
			if got := c.getLabels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("csvRow.getLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
