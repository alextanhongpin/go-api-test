package api

import "time"

type Health struct {
	ServiceName string    `json:"serviceName"`
	Version     string    `json:"version"`
	BuildAt     time.Time `json:"buildAt"`
	StartAt     time.Time `json:"startAt"`
	Uptime      string    `json:"uptime"`
	VCSRef      string    `json:"vcsRef"`
	VCSURL      string    `json:"vcsUrl"`
}
