package server

import (
	"bytes"
	"github.com/dchest/captcha"
	"log"
	"net/http"
	"path"
	"strings"
	"time"
)

type captchaHandler struct {
	imgWidth  int
	imgHeight int
}

func CaptchaServer(imgWidth, imgHeight int) http.Handler {
	//	log.Println("starting CaptchaServer")
	return &captchaHandler{imgWidth, imgHeight}
}

func (h *captchaHandler) serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool) error {
	log.Println("serve")

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		log.Println("png")
		w.Header().Set("Content-Type", "image/png")
		captcha.WriteImage(&content, id, h.imgWidth, h.imgHeight)
	case ".wav":
		log.Println("wav")
		w.Header().Set("Content-Type", "audio/x-wav")
		captcha.WriteAudio(&content, id, lang)
	default:
		log.Println("404")
		return captcha.ErrNotFound
	}

	if download {
		log.Println("download")
		w.Header().Set("Content-Type", "application/octet-stream")
	}

	//	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	b := bytes.NewReader(content.Bytes())
	log.Println("bytes NewReader ServeContent")
	http.ServeContent(w, r, id+ext, time.Time{}, b)
	return nil
}

func (h *captchaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("ServeHTTP")

	dir, file := path.Split(r.URL.Path)
	//	log.Println(r.URL.Path)

	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]

	//	log.Println(ext)
	//	log.Println(id)
	if ext == "" || id == "" {
		http.NotFound(w, r)
		return
	}
	if r.FormValue("reload") != "" {
		captcha.Reload(id)
	}
	lang := strings.ToLower(r.FormValue("lang"))
	download := path.Base(dir) == "download"
	if h.serve(w, r, id, ext, lang, download) == captcha.ErrNotFound {
		http.NotFound(w, r)
	}
	// Ignore other errors.
}
