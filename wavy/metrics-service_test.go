package wavy

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
)

func Test_metricsService_GetTotalListens(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	c := NewClient(ctx, hclog.NewNullLogger(), os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.MetricsService().GetTotalListens(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("metricsService.GetTotalListens() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.NotZero(t, got)
		})
	}
}

func Test_metricsService_GetTotalUsers(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	c := NewClient(ctx, hclog.NewNullLogger(), os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.MetricsService().GetTotalUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("metricsService.GetTotalUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.NotZero(t, got)
		})
	}
}

func Test_metricsService_GetUserListensLeaderboard(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	c := NewClient(ctx, hclog.NewNullLogger(), os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.MetricsService().GetUserListensLeaderboard(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("metricsService.GetUserListensLeaderboard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.NotNil(t, got)
			assert.NotZero(t, got[0].Count)
			assert.NotZero(t, got[0].Username)
			assert.NotZero(t, got[0].UserID)
		})
	}
}
