package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func downloadFile(url string, filepath string) error {
	// set timeout
	timeout := time.Duration(35 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	// set user agent
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")

	// get data
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// create file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// write body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}

		fpath := filepath.Join(dest, f.Name)
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("%s: illegal file path", fpath)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			os.MkdirAll(filepath.Dir(fpath), os.ModePerm)
			f, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
			f.Close()
			rc.Close()
		}
	}
	return nil
}

func safeSubstring(s string, start, end int) string {
	var startIndex int
	var endIndex int
	if start < 0 {
		startIndex = 0
	} else {
		startIndex = start
	}
	if end > len(s) {
		endIndex = len(s)
	} else {
		endIndex = end
	}
	// strip out trailing commas
	for endIndex > startIndex && s[endIndex-1] == ',' {
		endIndex--
	}

	// if all spaces, return empty string
	if strings.TrimSpace(s[startIndex:endIndex]) == "" {
		return ""
	}

	//DEBUG: fmt.Printf("startIndex: %d, endIndex: %d\n", startIndex, endIndex)
	return s[startIndex:endIndex]
}
