package upload

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/atotto/clipboard"
	"github.com/baoist/transporter/notify"
	"github.com/stacktic/dropbox"
)

func Upload(path string) (string, error) {
	return upload(path)
}

func upload(path string) (string, error) {
	tmpPath, tmpFile := buildTmpFile(path)

	entry, err := db.UploadFile(tmpPath, tmpFile, false, "")
	if err != nil {
		return "", errors.New("Unable to save temporary file in Dropbox.")
	}

	go createLink(entry)

	entry, err = db.UploadFile(path, tmpFile, true, "")
	if err != nil {
		return "", errors.New("Unable to upload created file.")
	}

	return "", nil
}

func createLink(entry *dropbox.Entry) {
	// build a shareable link in dropbox
	link, _ := db.Media(entry.Path)

	// copy URL to clipboard
	clipboard.WriteAll(link.URL)

	// create an OS notification
	notify.Notify("Transporter", "Uploaded, pasted to clipboard.")
}

func buildTmpFile(path string) (tmpPath, filename string) {
	filename = filepath.Base(path)
	tmpPath = fmt.Sprintf("/tmp/transporter/%s", filename)

	mode := int(0777)
	_ = ioutil.WriteFile(tmpPath, []byte{}, os.FileMode(mode))

	return tmpPath, filename
}
