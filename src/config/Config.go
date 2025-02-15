package config

import (
	"time"
)

const BoardSize = 6

// -ldflags "-X github.com/HuXin0817/dots-and-boxes/config.debug=true"
var debug string
var DEBUG = debug == "true"

const PlayerTimeOut = 90 * time.Second
