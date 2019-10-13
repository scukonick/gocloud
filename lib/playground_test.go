package lib

import (
	"testing"
)

func TestDefaultComputer_Run(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"", args{""}, "", true},
		{"", args{`package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
}
`}, "Hello, playground\n", false},
		{"", args{`package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Println("Hello, playground")
}
`}, "Hello, playground\nHello, playground\n", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewDefault(nil)
			got, err := c.Run(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Run() got = %v, want %v", got, tt.want)
			}
		})
	}
}
