package wavy

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/go-hclog"
)

// UserService Reference for accessing public user profiles.
// https://wavy.fm/developers/docs/v1beta/users
type UserService interface {
	// GetProfile
	// Retrieves the public profile of a wavy.fm user. Note that private profiles will not be returned at all by this endpoint, regardless of authorization scopes.
	GetProfile(ctx context.Context, uri UserURI) (*GetUserProfileResponse, error)
	// HistroyService this service gives access to the /history endpoints
	HistroyService(uri UserURI) UserHistoryService
}

type userService struct {
	c      *client
	logger hclog.Logger
}

func newUserService(c *client, logger hclog.Logger) UserService {
	subLogger := logger.Named("user-service")

	return &userService{
		logger: subLogger,
		c:      c,
	}
}

// GetProfile
// Retrieves the public profile of a wavy.fm user. Note that private profiles will not be returned at all by this endpoint, regardless of authorization scopes.
func (u *userService) GetProfile(ctx context.Context, uri UserURI) (*GetUserProfileResponse, error) {
	u.logger.Trace("fetching user profile")
	defer u.logger.Trace("finished fetching user profile")

	res, err := u.c.get(fmt.Sprintf("/users/%s", uri.String()))
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get user profile for %q: %w", u.logger.Name(), uri.String(), err)
	}

	var userProfile GetUserProfileResponse
	err = json.NewDecoder(res.Body).Decode(&userProfile)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to parse response body of user profile: %w", u.logger.Name(), err)
	}

	return &userProfile, nil
}

func (u *userService) HistroyService(uri UserURI) UserHistoryService {
	return newUserHistryService(uri, u.c, u.logger)
}

// UserURI Represents a data strcut to build proper UserURI. Only one of the fields should be set.
// UserURI is explaned on this page: https://wavy.fm/developers/docs/v1beta/overview#user-uris
type UserURI struct {
	Username  string
	UserID    string
	DiscordID string
}

func (u *UserURI) UnmarshalBinary(data []byte) error {
	dataString := string(data)
	pieces := strings.Split(dataString, ":")

	if len(pieces) < 3 {
		return fmt.Errorf("failed to parse UserURI: %q", dataString)
	}

	switch pieces[2] {
	case "id":
		u.UserID = pieces[3]
	case "username":
		u.Username = pieces[3]
	case "discord":
		u.DiscordID = pieces[3]
	default:
		return fmt.Errorf("failed to parse UserURI: %q", dataString)
	}

	return nil
}

// String returns a properly formatted uri depending on which field is set on the struct.
func (u *UserURI) String() string {
	if u.UserID != "" {
		return fmt.Sprintf("wavyfm:user:id:%s", u.UserID)
	}
	if u.Username != "" {
		return fmt.Sprintf("wavyfm:user:username:%s", u.Username)
	}
	if u.DiscordID != "" {
		return fmt.Sprintf("wavyfm:user:discord:%s", u.DiscordID)
	}

	return "<UserURI not set correcntly>"
}

// ParseUserURI takes a full uri string and parses it into a struct.
// The uri formats are explained at: https://wavy.fm/developers/docs/v1beta/overview#user-uris
func ParseUserURI(uri string) (*UserURI, error) {
	r := &UserURI{}
	err := r.UnmarshalBinary([]byte(uri))
	if err != nil {
		return nil, err
	}
	return r, nil
}

type GetUserProfileResponse struct {
	URI      string    `json:"uri"`
	ID       string    `json:"id"`
	Username string    `json:"username"`
	JoinTime time.Time `json:"join_time"`
	Profile  Profile   `json:"profile"`
}

type Spotify struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
}

type Discord struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
}

type Profile struct {
	URL         string  `json:"url"`
	Avatar      string  `json:"avatar"`
	AvatarSmall string  `json:"avatar_small"`
	Country     string  `json:"country"`
	Biography   string  `json:"biography"`
	Twitter     string  `json:"twitter"`
	Instagram   string  `json:"instagram"`
	Spotify     Spotify `json:"spotify"`
	Discord     Discord `json:"discord"`
}
