package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shawu21/test/router"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat:           "2006-01-02 15:04:05",
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		ForceColors:               true,
	})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	//log.SetLevel(log.WarnLevel)
}

func main() {
	r := gin.Default()
	router.Routers(r)
	r.Run()
}
