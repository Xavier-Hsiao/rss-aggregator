package app

import (
	"github.com/Xavier-Hsiao/rss-aggregator/internal/config"
	"github.com/Xavier-Hsiao/rss-aggregator/internal/datbase"
)

type State struct {
	Config *config.Config
	DB     *datbase.Queries
}
