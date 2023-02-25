package reporting

import (
	"reflect"
	"testing"
)

func Test_parseSeparatedString(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "comma",
			args: args{
				s:   "a,b,c",
				sep: ",",
			},
			want:    []string{"a", "b", "c"},
			wantErr: false,
		},
		{
			name: "trimmed comma",
			args: args{
				s:   "a, b, c",
				sep: ",",
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "excessive whitespace",
			args: args{
				s:   "     			a,   b,    			c  	",
				sep: ",",
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "empty",
			args: args{
				s:   "",
				sep: ",",
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseSeparatedString(tt.args.s, tt.args.sep)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseSeparatedString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSeparatedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
