package wavy

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
)

func ExampleNewClient() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := NewClient(ctx, hclog.NewNullLogger(), os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	profile, err := c.UserService().GetProfile(ctx, UserURI{Username: "OGKevin"})
	if err != nil {
		panic(err)
	}

	fmt.Print(profile.URI)
}

func TestNewClient(t *testing.T) {
	type args struct {
		ctx          context.Context
		logger       hclog.Logger
		clientID     string
		clientSecret string
	}
	tests := []struct {
		name    string
		args    args
		want    Client
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:          context.Background(),
				logger:       hclog.New(&hclog.LoggerOptions{Level: hclog.Trace}),
				clientID:     os.Getenv("CLIENT_ID"),
				clientSecret: os.Getenv("CLIENT_SECRET"),
			},
			wantErr: false,
		},
		{
			name: "without logger",
			args: args{
				ctx:          context.Background(),
				logger:       nil,
				clientID:     "",
				clientSecret: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewClient(tt.args.ctx, tt.args.logger, tt.args.clientID, tt.args.clientSecret)
			assert.NotNil(t, got)
		})
	}
}
