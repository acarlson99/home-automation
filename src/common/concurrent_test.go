package common

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestConcurrentAggregateErrorFn(t *testing.T) {
	type args struct {
		f    func(int) error
		args []int
	}
	tests := []struct {
		name    string
		args    args
		wantErr int
	}{
		{
			name: "all good",
			args: args{
				f:    func(i int) error { return nil },
				args: []int{1, 2, 3, 4, 5},
			},
			wantErr: 0,
		},
		{
			name: "all bad",
			args: args{
				f:    func(i int) error { return fmt.Errorf("bad oh no") },
				args: []int{1, 2, 3, 4, 5},
			},
			wantErr: 5,
		},
		{
			name: "half fail",
			args: args{
				f: func(i int) error {
					if i >= 3 {
						return fmt.Errorf("bad oh no")
					}
					return nil
				},
				args: []int{1, 2, 3, 4, 5},
			},
			wantErr: 3,
		},
		{
			name: "half fail with sleep should not skip",
			args: args{
				f: func(i int) error {
					if i >= 3 {
						time.Sleep(time.Second * 3)
						return fmt.Errorf("bad oh no")
					}
					return nil
				},
				args: []int{1, 2, 3, 4, 5},
			},
			wantErr: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ConcurrentAggregateErrorFn(tt.args.f, tt.args.args...); err != nil && tt.wantErr != 0 {
				n := len(strings.Split(err.Error(), "\n"))
				if n != tt.wantErr {
					t.Errorf("ConcurrentlyAggregateErrors() errors = %v, wantErr %v", n, tt.wantErr)
				}
			}
		})
	}
}
