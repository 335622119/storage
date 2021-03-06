package main

import (
	"dxkite.cn/storage"
	"flag"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	var save = flag.String("path", "", "download save path")

	var install = flag.Bool("install", false, "install")
	var uninstall = flag.Bool("uninstall", false, "uninstall")

	var uncheck = flag.Bool("uncheck", false, "uncheck hash after downloaded")

	var help = flag.Bool("help", false, "print help info")
	var num = flag.Int("threads", 20, "max download threads")
	var retry = flag.Int("retry", 20, "max retry when error")

	flag.Parse()
	p, _ := filepath.Abs(os.Args[0])

	if *install {
		storage.Install(p)
		return
	}

	if *uninstall {
		storage.Uninstall(p)
		return
	}

	if *help || flag.NArg() < 1 {
		flag.Usage()
		return
	}

	name := flag.Arg(0)
	if storage.FileExist(name) && strings.HasSuffix(name, storage.EXT_META) {
		if len(*save) == 0 {
			p, _ := filepath.Abs(name)
			*save = filepath.Dir(p)
		}
		storage.Download(name, *save, *uncheck == false, *num, *retry)
	} else {
		if len(*save) == 0 {
			pp := filepath.Dir(p)
			*save = path.Join(pp, "Download")
			_ = os.MkdirAll(*save, os.ModePerm)
		}
		storage.Download(name, *save, *uncheck == false, *num, *retry)
	}
}
