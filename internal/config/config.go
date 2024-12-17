package config

import "time"

const (
	AsanaAPIBase        = "https://app.asana.com/api/1.0"
	OutputFolder        = "./output"
	RequestsPerSec      = 5
	WorkerCount         = 5
	FetchIntervalFast   = 30 * time.Second
	FetchIntervalSlow   = 5 * time.Minute
	FetchIntervalFastTest = 5 * time.Second
	FetchIntervalSlowTest = 10 * time.Second
	AsanaAccessToken    = "Bearer 2/1208996646709526/1208996863373258:10066150db2c3920f16f483c00c259d5"
)
