package common

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

type testFnCounter struct {
	n int
}

func (fnc *testFnCounter) callF(n int) error {
	fnc.n += 1
	if n < 10 {
		return fmt.Errorf("ERR")
	}
	return nil
}

func TestConcurrentAggregateErrorFn(t *testing.T) {
	for i := 0; i < 20; i++ {
		t.Run(fmt.Sprintf("try-force-concurrency-fail-%02d", i), func(t *testing.T) {
			t.Parallel()
			tfnc := &testFnCounter{n: 0}
			type args struct {
				f    func(int) error
				args []int
			}
			tests := []struct {
				name      string
				args      args
				wantErr   int
				useStruct bool
			}{
				{
					name: "ensure all fns run",
					args: args{
						f:    tfnc.callF,
						args: []int{1, 2, 3, 4, 5},
					},
					wantErr:   5,
					useStruct: true,
				},
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
								time.Sleep(time.Second * time.Duration(float32(i)*0.2))
								return fmt.Errorf("bad oh no")
							}
							return nil
						},
						args: []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5},
					},
					wantErr: 12,
				},
				{
					name: "all fail with sleep should not skip 2",
					args: args{
						f: func(i int) error {
							time.Sleep(time.Second * time.Duration(float32(i)*0.5))
							return fmt.Errorf("bad oh no")
						},
						args: []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5},
					},
					wantErr: 20,
				},
			}
			for _, tt := range tests {
				tfnc.n = 0
				t.Run(tt.name, func(t *testing.T) {
					if err := ConcurrentAggregateErrorFn(tt.args.f, tt.args.args...); err != nil {
						n := len(strings.Split(err.Error(), "\n"))
						if n != tt.wantErr {
							t.Errorf("ConcurrentlyAggregateErrors() errors = %v, wantErr %v", n, tt.wantErr)
						}
					}
					if tt.useStruct {
						if got, want := tfnc.n, len(tt.args.args); got != want {
							t.Errorf("Not all functions were executed: got %d want %d", got, want)
						}
					}
				})
			}
		})
	}
}
