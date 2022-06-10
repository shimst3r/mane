package mane

import (
	"testing"
)

func Test_fromEU(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{"test 1", args{"1,00"}, 100, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fromEU(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("fromEU() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("fromEU() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	type args struct {
		input       string
		source      Locale
		destination Locale
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"EU to US 1", args{"1,0", "EU", "US"}, "1.0", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Convert(tt.args.input, tt.args.source, tt.args.destination)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
