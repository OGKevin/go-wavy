package wavy

import (
	"context"
	"os"
	"testing"

	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
)

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
