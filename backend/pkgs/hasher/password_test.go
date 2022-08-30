package hasher

import "testing"

func TestHashPassword(t *testing.T) {
	t.Parallel()
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "letters_and_numbers",
			args: args{
				password: "password123456788",
			},
		},
		{
			name: "letters_number_and_special",
			args: args{
				password: "!2afj3214pofajip3142j;fa",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !CheckPasswordHash(tt.args.password, got) {
				t.Errorf("CheckPasswordHash() failed to validate password=%v against hash=%v", tt.args.password, got)
			}
		})
	}
}
