package repo

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAssetID_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		aid     AssetID
		want    []byte
		wantErr bool
	}{
		{
			name: "basic test",
			aid:  123,
			want: []byte(`"000-123"`),
		},
		{
			name: "zero test",
			aid:  0,
			want: []byte(`""`),
		},
		{
			name: "large int",
			aid:  123456789,
			want: []byte(`"123-456789"`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.aid.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("AssetID.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssetID.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssetID_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		aid     *AssetID
		args    args
		want    AssetID
		wantErr bool
	}{
		{
			name: "basic test",
			aid:  new(AssetID),
			want: 123,
			args: args{
				data: []byte(`{"AssetID":"000123"}`),
			},
		},
		{
			name: "dashed format",
			aid:  new(AssetID),
			want: 123,
			args: args{
				data: []byte(`{"AssetID":"000-123"}`),
			},
		},
		{
			name: "no leading zeros",
			aid:  new(AssetID),
			want: 123,
			args: args{
				data: []byte(`{"AssetID":"123"}`),
			},
		},
		{
			name: "trailing zeros",
			aid:  new(AssetID),
			want: 123000,
			args: args{
				data: []byte(`{"AssetID":"000123000"}`),
			},
		},
		{
			name: "large int",
			aid:  new(AssetID),
			want: 123456789,
			args: args{
				data: []byte(`{"AssetID":"123456789"}`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := struct {
				AssetID AssetID `json:"AssetID"`
			}{}

			err := json.Unmarshal(tt.args.data, &st)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssetID.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if st.AssetID != tt.want {
				t.Errorf("AssetID.UnmarshalJSON() = %v, want %v", st.AssetID, tt.want)
			}
		})
	}
}
