package geoip2

import (
	"log"
	"net"
	"testing"
	"time"
	 "github.com/apilayer/freegeoip"
)
func TestFindIP(t *testing.T) {
	updateInterval := 24 * time.Hour
	maxRetryInterval := time.Hour
	db, err := OpenURL(MaxMindDB, updateInterval, maxRetryInterval)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	select {
	case <-db.NotifyOpen():
		// Wait for the db to be downloaded.
	case err := <-db.NotifyError():
		log.Fatal(err)
	}
	var result customQuery
	err = db.Lookup(net.ParseIP("8.8.8.8"), &result)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v", result)
}