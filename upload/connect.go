package upload

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/stacktic/dropbox"
)

var (
	appKey    = os.Getenv("TRANSPORTER_APP_KEY")
	appSecret = os.Getenv("TRANSPORTER_APP_SECRET")
	appToken  = os.Getenv("TRANSPORTER_APP_TOKEN")
	db        *dropbox.Dropbox
)

func Connect() {
	if err := connect(); err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to Dropbox.")
}

func connect() error {
	if appKey == "" || appSecret == "" {
		return errors.New(fmt.Sprintf("Unable to connect to dropbox. Env vars not set.\r\n"+
			"TRANSPORTER_APP_KEY: %s\r\n"+
			"TRANSPORTER_APP_SECRET: %s\r\n"+
			"TRANSPORTER_APP_TOKEN: %s", appKey, appSecret, appToken))
	}

	buildTmpDirectory()

	db = dropbox.NewDropbox()
	db.SetAppInfo(appKey, appSecret)
	db.SetAccessToken(appToken)

	return nil
}

func buildTmpDirectory() {
	mode := int(0777)
	_ = os.Mkdir("/tmp/transporter", os.FileMode(mode))
}
