package upload

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/atotto/clipboard"
	"github.com/baoist/transporter/notify"
)

func Upload(path string) (string, error) {
	return upload(path)
}

func upload(path string) (string, error) {
	tmpPath, tmpFile := buildTmpFile(path)

	entry, err := db.UploadFile(tmpPath, tmpFile, false, "")
	if err != nil {
		return "", errors.New("Unable to create temporary file.")
	}

	link, err := db.Media(entry.Path)
	if err != nil {
		return "", errors.New("Unable to generate link")
	}

	clipboard.WriteAll(link.URL)
	notify.Notify("Transporter", "Uploaded and pasted to clipboard.")

	entry, err = db.UploadFile(path, tmpFile, true, "")
	if err != nil {
		return "", errors.New("Unable to upload created file.")
	}

	return "", nil
}

func buildTmpFile(path string) (tmpPath, filename string) {
	filename = filepath.Base(path)
	tmpPath = fmt.Sprintf("/tmp/transporter/%s", filename)

	mode := int(0777)
	_ = ioutil.WriteFile(tmpPath, []byte{}, os.FileMode(mode))

	return tmpPath, filename
}
