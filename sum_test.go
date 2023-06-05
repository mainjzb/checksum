package utils

import (
	"testing"
)

func TestBBC(t *testing.T) {
	tests := []struct {
		name string
		args []byte
		want uint8
	}{
		{"success 1", []byte{0x01, 0x02}, 0x03},
		{"success 2", []byte{0x01, 0x02, 0x03}, 0x00},
		{"success 3", []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06}, 0x07},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BBC(tt.args); got != tt.want {
				t.Errorf("BBC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		args []byte
		want uint16
	}{
		{"success 1", []byte{0x01, 0x02}, 0x03},
		{"success 2", []byte{0x01, 0x02, 0x03}, 0x06},
		{"success 3", []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06}, 0x15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum2(tt.args); got != tt.want {
				t.Errorf("BBC() = %v, want %v", got, tt.want)
			}
		})
	}
}
