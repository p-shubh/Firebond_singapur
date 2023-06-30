package response

import (
	log "github.com/sirupsen/logrus"
)

// Version Build Version
var Version = "1.3"

// StandardFields for logger
var StandardFields = log.Fields{
	"hostname": "platform",
	"appname":  "Platform",
}
