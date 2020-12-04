package helper

import (
	"errors"
	"testing"
)

func TestCheck(t *testing.T) {
	type args struct {
		err error
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "With error",
			args: args{
				err: errors.New("Some Error"),
				msg: "Some error occured",
			},
		}, {
			name: "Without error",
			args: args{
				err: nil,
				msg: "Some error occured",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Check(tt.args.err, tt.args.msg)
		})
	}
}
