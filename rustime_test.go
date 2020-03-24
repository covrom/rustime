package rustime

import (
	"testing"
	"time"
)

func TestFormatTimeRu(t *testing.T) {
	type args struct {
		t      time.Time
		fmtstr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				t:      time.Date(2019, 10, 19, 5, 35, 44, int(550*time.Millisecond), time.UTC),
				fmtstr: "дд ММММ гггг чч:мм:сс.ссс",
			},
			want: "19 октября 2019 05:35:44.550",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatTimeRu(tt.args.t, tt.args.fmtstr); got != tt.want {
				t.Errorf("FormatTimeRu() = %v, want %v", got, tt.want)
			}
		})
	}
}
