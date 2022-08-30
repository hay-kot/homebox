package types

import (
	"testing"
)

func TestUserCreate_Validate(t *testing.T) {
	type fields struct {
		Name        string
		Email       string
		Password    string
		IsSuperuser bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "no_name",
			fields: fields{
				Name:        "",
				Email:       "",
				Password:    "",
				IsSuperuser: false,
			},
			wantErr: true,
		},
		{
			name: "no_email",
			fields: fields{
				Name:        "test",
				Email:       "",
				Password:    "",
				IsSuperuser: false,
			},
			wantErr: true,
		},
		{
			name: "valid",
			fields: fields{
				Name:        "test",
				Email:       "test@email.com",
				Password:    "mypassword",
				IsSuperuser: false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserCreate{
				Name:        tt.fields.Name,
				Email:       tt.fields.Email,
				Password:    tt.fields.Password,
				IsSuperuser: tt.fields.IsSuperuser,
			}
			if err := u.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("UserCreate.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
