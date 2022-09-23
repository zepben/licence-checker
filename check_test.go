package main

import (
	"testing"
)

func TestIsValidLicense(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "Zepben copy", args: args{"Zepben-2023.txt"}, want: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsValidLicense(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsValidLicense() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsValidLicense() got = %v, want %v", got, tt.want)
			}
		})
	}
}
