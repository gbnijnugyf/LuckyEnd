package controller

import (
	"strconv"
	"testing"

	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

func TestLog(t *testing.T) {
	_, err := strconv.Atoi("u")
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat:           "2006-01-02 15:04:05",
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		ForceColors:               true,
	})
	log.Errorf("cannot convert type into int err :%+v", errors.WithStack(err))
}
