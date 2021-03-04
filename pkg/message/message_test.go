package message

import (
	"testing"
	"time"
)

func TestNiceUptime(t *testing.T) {
	type args struct {
		start time.Time
		now   time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TwoMin", args: args{start: time.Unix(30, 0), now: time.Unix(90, 0)}, want: "1m0s"},
		{name: "Hundred Days", args: args{start: time.Unix(0, 0), now: time.Unix(60*60*24*100, 0)}, want: "2400h0m0s"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NiceUptime(tt.args.start, tt.args.now); got != tt.want {
				t.Errorf("NiceUptime() = %v, want %v", got, tt.want)
			}
		})
	}
}
