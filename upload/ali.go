package upload

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

//阿里
type AliResp struct {
	FsUrl string `json:"fs_url"`
	Code  string `json:"code"`
	Size  string `json:"size"`
	Width string `json:"width"`
	Url   string `json:"url"`
	Hash  string `json:"hash"`
}

const ALI = "ali"

func init() {
	// 注册阿里文件图床
	Register(ALI, func(config Config) Uploader {
		return &AliUploader{}
	})
}

type AliUploader struct {
}

func (*AliUploader) Upload(object *FileObject) (*Result, error) {
	url := "https://kfupload.alibaba.com/mupload"
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	if fw, e := w.CreateFormField("scene"); e == nil && fw != nil {
		if _, er := fw.Write([]byte("aeMessageCenterV2ImageRule")); er != nil {
			return nil, er
		}
	}
	if fw, e := w.CreateFormField("name"); e == nil && fw != nil {
		if _, er := fw.Write([]byte(object.Name)); er != nil {
			return nil, er
		}
	}
	if fw, e := w.CreateFormFile("file", object.Name); e == nil && fw != nil {
		if _, er := fw.Write(object.Data); er != nil {
			return nil, er
		}
	}
	if er := w.Close(); er != nil {
		return nil, er
	}

	req, _ := http.NewRequest(http.MethodPost, url, &b)
	req.Header.Set("Content-Type", w.FormDataContentType())

	res, er := Client.Do(req)
	if er != nil {
		return nil, er
	}

	defer func() { _ = res.Body.Close() }()

	body, rer := ioutil.ReadAll(res.Body)
	if rer != nil {
		return nil, rer
	}
	resp := new(AliResp)
	if er := json.Unmarshal(body, resp); er == nil {
		if resp.Code != "0" {
			return nil, errors.New("ali upload error, code:" + resp.Code)
		}
		return &Result{
			Url: resp.Url,
			Raw: body,
		}, nil
	} else {
		return nil, er
	}
}
