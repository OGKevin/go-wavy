package wavy

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-hclog"
)

type UserHistoryService interface {
	// GetHistryStats
	// Retrieves some statistics about the user's history. Note that private profiles will not be returned at all by this endpoint, regardless of authorization scopes.
	GetStats() (*GetHistroyStatsResponse, error)
	// GetCurrent
	// Retrieves the song, album, and artist(s) the user is currently listening to. Note that private profiles will not be returned at all by this endpoint, regardless of authorization scopes.
	GetCurrent() (*GetCurrentResponse, error)
	// GetRecent
	// Retrieves the most recent listens recorded by the user. Note that private profiles will not be returned at all by this endpoint, regardless of authorization scopes.
	GetRecent() (*GetRecentResponse, error)
}

type userHistroyService struct {
	userUri UserURI

	c      *client
	logger hclog.Logger
}

func newUserHistryService(uri UserURI, c *client, logger hclog.Logger) UserHistoryService {
	subLogger := logger.Named("histroy-service")

	return &userHistroyService{
		userUri: uri,
		c:       c,
		logger:  subLogger,
	}
}

func (u *userHistroyService) buildUrl(path string) string {
	return fmt.Sprintf("/user/%s/histroy%s", u.userUri, path)
}

// GetStats
// Retrieves some statistics about the user's history. Note that private profiles will not be returned at all by this endpoint, regardless of authorization scopes.
func (u *userHistroyService) GetStats() (*GetHistroyStatsResponse, error) {
	u.logger.Trace("fetching stats")
	defer u.logger.Trace("finished fetching stats")

	res, err := u.c.get(u.buildUrl("/stats"))
	if err != nil {
		return nil, fmt.Errorf("%s: failed to fetch stats for %q: %w", u.logger.Name(), u.userUri, err)
	}

	var statsRes GetHistroyStatsResponse
	err = json.NewDecoder(res.Body).Decode(&statsRes)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to parse response body for histroy stats %w", u.logger.Name(), err)
	}

	return &statsRes, nil
}

// GetCurrent
// Retrieves the song, album, and artist(s) the user is currently listening to. Note that private profiles will not be returned at all by this endpoint, regardless of authorization scopes.
func (u *userHistroyService) GetCurrent() (*GetCurrentResponse, error) {
	u.logger.Trace("fetching current")
	defer u.logger.Trace("finished fetching current")

	res, err := u.c.get(u.buildUrl("/current"))
	if err != nil {
		return nil, fmt.Errorf("%s: failed to fetch current for %q: %w", u.logger.Name(), u.userUri, err)
	}

	var currentRes GetCurrentResponse
	err = json.NewDecoder(res.Body).Decode(&currentRes)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to parse response body for current %w", u.logger.Name(), err)
	}

	return &currentRes, nil
}

// GetRecent
// Retrieves the most recent listens recorded by the user. Note that private profiles will not be returned at all by this endpoint, regardless of authorization scopes.
func (u *userHistroyService) GetRecent() (*GetRecentResponse, error) {
	panic("not implemented") // TODO: Implement
	u.logger.Trace("fetching recent")
	defer u.logger.Trace("finished fetching recent")

	res, err := u.c.get(u.buildUrl("/recent"))
	if err != nil {
		return nil, fmt.Errorf("%s: failed to fetch recent for %q: %w", u.logger.Name(), u.userUri, err)
	}

	var recentRes GetRecentResponse
	err = json.NewDecoder(res.Body).Decode(&recentRes)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to parse response body for recent %w", u.logger.Name(), err)
	}

	return &recentRes, nil
}

type GetHistroyStatsResponse struct {
	TotalListens int `json:"total_listens"`
	TotalArtists int `json:"total_artists"`
}

type GetCurrentResponse struct {
	Item CurrentPlayingItem `json:"item"`
}

type Song struct {
	Source    string `json:"source"`
	SourceURL string `json:"source_url"`
	Name      string `json:"name"`
}

type Album struct {
	Source    string `json:"source"`
	SourceURL string `json:"source_url"`
	Name      string `json:"name"`
	ArtURL    string `json:"art_url"`
}

type Artists struct {
	Source    string `json:"source"`
	SourceURL string `json:"source_url"`
	Name      string `json:"name"`
}

type CurrentPlayingItem struct {
	Local   bool      `json:"local"`
	Song    Song      `json:"song"`
	Album   Album     `json:"album"`
	Artists []Artists `json:"artists"`
}

type GetRecentResponse struct {
	Items []Item `json:"items"`
}

type Item struct {
	CurrentPlayingItem
	// Local   bool      `json:"local"`
	Date   time.Time `json:"date"`
	PlayID string    `json:"play_id"`
	// Song    Song      `json:"song"`
	// Album   Album     `json:"album"`
	// Artists []Artists `json:"artists"`
}
