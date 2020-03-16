package main

import (
	"dxkite.cn/go-storage/src/client"
	"encoding/hex"
	"flag"
	"log"
	"time"
)

func Upload(addr, path string) {
	u := client.NewUploader(addr, time.Second*100)
	if er := u.UploadFile(path); er != nil {
		log.Fatal("upload error:", er)
	}
	log.Println("upload success")
}

func UploadLocal(addr, path string) {
	u := client.NewLocalUploader(2*1024*1024, "ali")
	if er := u.UploadFile(path); er != nil {
		log.Fatal("upload error:", er)
	}
	log.Println("upload success")
}

func Download(addr, info, path string) {
	h, _ := hex.DecodeString(info)
	d := client.NewRemoteDownloader(addr, h)
	if er := d.DownloadToFile(path); er != nil {
		log.Fatal(er)
	}
	log.Println("download success")
}

func DownloadMeta(meta, path string) {
	d := client.NewMetaDownloader(meta)
	if er := d.DownloadToFile(path); er != nil {
		log.Fatal(er)
	}
	log.Println("download success")
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	var addr = flag.String("addr", "127.0.0.1:20214", "listening")
	var path = flag.String("path", "./", "download to path")
	var meta = flag.Bool("meta", false, "use meta file")
	var help = flag.Bool("help", false, "print help info")

	flag.Parse()
	if *help || flag.NArg() < 1 {
		flag.Usage()
		return
	}

	name := flag.Arg(0)

	if client.FileExist(name) {
		if *meta {
			DownloadMeta(name, *path)
		} else {
			UploadLocal(*addr, name)
		}
	} else {
		Download(*addr, name, *path)
	}
}
