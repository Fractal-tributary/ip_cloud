package geoip2

import (
	"fmt"
	"github.com/Fractal-tributary/freegeoip"
	"log"
	"net"
	"os"
	"testing"
	"time"
)

var MaxMindDB = "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City&license_key=R9bpXwUb0nfuavnC&suffix=tar.gz"

func TestFindIP(t *testing.T) {
	updateInterval := 24 * time.Hour
	maxRetryInterval := time.Hour
	fmt.Print( os.TempDir())
	db, err := freegeoip.OpenURL(MaxMindDB, updateInterval, maxRetryInterval)
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
	var result freegeoip.DefaultQuery
	err = db.Lookup(net.ParseIP("8.8.8.8"), &result)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v", result)
}

func TestFindIP2(t *testing.T) {
	db, err := freegeoip.Open("GeoLite2-City.mmdb.gz")
	fmt.Print(os.TempDir())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var result freegeoip.DefaultQuery
	err = db.Lookup(net.ParseIP("8.8.8.8"), &result)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v", result)
}