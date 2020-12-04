package helper

import "testing"

func TestPublish(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "WrongPath",
			args: args{
				path: "./input.json",
			},
			want: false,
		}, {
			name: "CorrectPath",
			args: args{
				path: "../../input.json",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Publish(tt.args.path); got != tt.want {
				t.Errorf("Publish() = %v, want %v", got, tt.want)
			}
		})
	}
}
