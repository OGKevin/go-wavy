package wavy

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
)

func Test_userService_GetProfile(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	c := NewClient(ctx, hclog.NewNullLogger(), os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))

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
			got, err := c.UserService().GetProfile(tt.args.ctx, tt.args.userURI)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.GetProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.NotZero(t, got)
			assert.NotZero(t, got.ID)
		})
	}
}

func Test_userService_ParseUserURI(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Empty string",
			args: args{
				uri: "",
			},
			wantErr: true,
		},
		{
			name: "Username",
			args: args{
				uri: "wavyfm:user:username:User",
			},
			wantErr: false,
		},
		{
			name: "Discord",
			args: args{
				uri: "wavyfm:user:discord:Discord",
			},
			wantErr: false,
		},

		{
			name: "id",
			args: args{
				uri: "wavyfm:user:id:uuid",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseUserURI(tt.args.uri)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUserURI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			assert.NotZero(t, got)
			assert.NotContains(t, got.String(), "UserURI not set correcntly")
			assert.Equal(t, tt.args.uri, got.String())
		})
	}
}
