package main

import (
	"flag"

	"github.com/baoist/transporter/upload"
	"github.com/baoist/transporter/watch"
)

var dir string

func main() {
	flag.StringVar(&dir, "dir", "/tmp", "Directory for transporter to watch")
	flag.Parse()

	upload.Connect()
	watch.Watch(dir, upload.Upload)
}
