package services

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//go:embed .testdata/import.csv
var CSVData_Comma []byte

//go:embed .testdata/import.tsv
var CSVData_Tab []byte

func loadcsv() [][]string {
	reader := csv.NewReader(bytes.NewReader(CSVData_Comma))

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

		assert.Equal(t, expected, entity.Item.PurchaseTime.Time(), fmt.Sprintf("Failed on row %d", i))
		assert.Equal(t, expected, entity.Item.WarrantyExpires.Time(), fmt.Sprintf("Failed on row %d", i))
		assert.Equal(t, expected, entity.Item.SoldTime.Time(), fmt.Sprintf("Failed on row %d", i))
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

func Test_determineSeparator(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    rune
		wantErr bool
	}{
		{
			name: "comma",
			args: args{
				data: CSVData_Comma,
			},
			want:    ',',
			wantErr: false,
		},
		{
			name: "tab",
			args: args{
				data: CSVData_Tab,
			},
			want:    '\t',
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				data: []byte("a;b;c"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := determineSeparator(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("determineSeparator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("determineSeparator() = %v, want %v", got, tt.want)
			}
		})
	}
}
