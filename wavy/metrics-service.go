package wavy

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/hashicorp/go-hclog"
)

// MetricsService
// Reference for accessing global wavy.fm metrics.
// https://wavy.fm/developers/docs/v1beta/metrics
type MetricsService interface {
	// GetTotalListens
	// Retrieves the total amount of listens recorded on wavy.fm. Note that this value is cached for a few seconds.
	GetTotalListens(ctx context.Context) (int, error)
	// GetTotalUsers
	// Retrieves the total amount of registered users on wavy.fm. Note that this value is cached for a few seconds.
	GetTotalUsers(ctx context.Context) (int, error)
	// GetUserListensLeaderboard
	// Retrieves the leaderboard of the top 10 users by listen count. Note that this endpoint is cached for a few minutes.
	GetUserListensLeaderboard(ctx context.Context) (UserListensLeaderboardResponse, error)
}

type metricsService struct {
	c      *client
	logger hclog.Logger
}

func newMetricsService(c *client, logger hclog.Logger) MetricsService {
	logger = logger.Named("metrics-service")

	return &metricsService{
		logger: logger,
		c:      c,
	}
}

// GetTotalListens
// Retrieves the total amount of listens recorded on wavy.fm. Note that this value is cached for a few seconds.
func (m *metricsService) GetTotalListens(ctx context.Context) (int, error) {
	m.logger.Trace("fetching total listens")
	defer m.logger.Trace("finished fetching total listens")

	res, err := m.c.get("/metrics/total-listens")
	if err != nil {
		return 0, fmt.Errorf("%s: failed to request total listens: %w", m.logger.Name(), err)
	}

	rawBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to parse response body: %w", m.logger.Name(), err)
	}

	totalListens, err := strconv.Atoi(string(rawBody))
	if err != nil {
		return 0, fmt.Errorf("%s: failed to parse response body to int: %w", m.logger.Name(), err)
	}

	return totalListens, nil
}

// GetTotalUsers
// Retrieves the total amount of registered users on wavy.fm. Note that this value is cached for a few seconds.
func (m *metricsService) GetTotalUsers(ctx context.Context) (int, error) {
	m.logger.Trace("fetching total users")
	defer m.logger.Trace("finished fetching total users")

	res, err := m.c.get("/metrics/total-users")
	if err != nil {
		return 0, fmt.Errorf("%s: failed to request total users: %w", m.logger.Name(), err)
	}

	rawBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to parse response body: %w", m.logger.Name(), err)
	}

	totalUsers, err := strconv.Atoi(string(rawBody))
	if err != nil {
		return 0, fmt.Errorf("%s: failed to parse response body to int: %w", m.logger.Name(), err)
	}

	return totalUsers, nil
}

// GetUserListensLeaderboard
// Retrieves the leaderboard of the top 10 users by listen count. Note that this endpoint is cached for a few minutes.
func (m *metricsService) GetUserListensLeaderboard(ctx context.Context) (UserListensLeaderboardResponse, error) {
	m.logger.Trace("fetching user listen leaderboards")
	defer m.logger.Trace("finished fetching user listen leaderboards")

	res, err := m.c.get("/metrics/user-listens-leaderboard")
	if err != nil {
		return nil, fmt.Errorf("%s: failed to request total users: %w", m.logger.Name(), err)
	}

	var leaderBoard UserListensLeaderboardResponse
	err = json.NewDecoder(res.Body).Decode(&leaderBoard)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to parse response body to struct: %w", m.logger.Name(), err)
	}

	return leaderBoard, nil
}

type UserListensLeaderboardResponse []struct {
	Count    int    `json:"count"`
	Username string `json:"username"`
	UserID   string `json:"user_id"`
}
