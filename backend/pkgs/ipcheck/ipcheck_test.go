package ipcheck

import (
	"testing"
)

func Test_ValidateAgainstList(t *testing.T) {
	tests := []struct {
		name string
		ip   string
		list string
		want bool
	}{
		{
			name: "IPv4 matching the list",
			ip:   "192.168.1.1",
			list: "192.168.11.0/24,192.168.1.0/24",
			want: true,
		}, {
			name: "IPv4 with exact match",
			ip:   "192.168.2.2",
			list: "192.168.2.2/32,192.168.0.0/24",
			want: true,
		}, {
			name: "IPv4 with no match",
			ip:   "192.168.3.3",
			list: "192.168.0.0/24,192.168.2.0/24",
			want: false,
		}, {
			name: "IPv6 matching the list",
			ip:   "1111:1111:1111:1111:1111:1111:1111:1111",
			list: "1111:1111:1111:1111::/64,2222:2222:2222:2222::/64",
			want: true,
		}, {
			name: "IPv6 with exact match",
			ip:   "2222:2222:2222:2222:2222:2222:2222:2222",
			list: "1111:1111:1111:1111::/64,2222:2222:2222:2222:2222:2222:2222:2222/128",
			want: true,
		}, {
			name: "IPv6 with no match",
			ip:   "3333:3333:3333:3333:3333:3333:3333:3333",
			list: "3333:3333:3333:3333:3333:3333:3333:4444/128,4444::/32",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateAgainstList(tt.ip, tt.list); got != tt.want {
				t.Errorf("ValidateAgainstList() = %v, want %v", got, tt.want)
			}
		})
	}
}
