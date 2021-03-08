package wavy

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
)

func Test_userHistroyService_GetStats(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	c := NewClient(ctx, hclog.NewNullLogger(), os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))

	type fields struct {
		c      *client
		logger hclog.Logger
	}
	type args struct {
		ctx     context.Context
		userURI UserURI
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
				userURI: UserURI{
					Username: "OGKevin",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.UserService().HistroyService(tt.args.userURI).GetStats(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("HistroyService.GetStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.NotZero(t, got)
			assert.NotZero(t, got.TotalListens)
			assert.NotZero(t, got.TotalArtists)
		})
	}
}

func Test_userHistroyService_GetCurrent(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	c := NewClient(ctx, hclog.NewNullLogger(), os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))

	type fields struct {
		c      *client
		logger hclog.Logger
	}
	type args struct {
		ctx     context.Context
		userURI UserURI
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
				userURI: UserURI{
					Username: "OGKevin",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.UserService().HistroyService(tt.args.userURI).GetCurrent(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("HistroyService.GetStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.NotZero(t, got)
		})
	}
}

func Test_userHistroyService_GetRecent(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	c := NewClient(ctx, hclog.NewNullLogger(), os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))

	type fields struct {
		c      *client
		logger hclog.Logger
	}
	type args struct {
		ctx     context.Context
		userURI UserURI
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
				userURI: UserURI{
					Username: "OGKevin",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.UserService().HistroyService(tt.args.userURI).GetRecent(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("HistroyService.GetStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.NotZero(t, got)
			for _, item := range got.Items {
				assert.NotZero(t, item.Song.Name)
			}
		})
	}
}
