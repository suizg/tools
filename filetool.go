package main

import (
	"bufio"
	"compress/gzip"
	"io"
	"net/http"
	"os"
)

func ReadLine(r *bufio.Reader) (string, error) {

	line, isprefix, err := r.ReadLine()
	for isprefix && err == nil {
		var bs []byte
		bs, isprefix, err = r.ReadLine()
		line = append(line, bs...)
	}
	return string(line), err
}

func pathExist(path string) (bool, error) {

	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func PathExist(path string) bool {

	b, _ := pathExist(path)
	if b == true {
		return true
	}
	return false
}

func RemoveFile(path string) error {

	err := os.Remove(path)
	return err
}

func DownFile(filename string, url string) (int64, error) {

	res, err := http.Get(url)
	if err != nil {
		return 0, nil
	}
	f, err := os.Create(filename)
	if err != nil {
		return 0, nil
	}

	sz, err := io.Copy(f, res.Body)
	return sz, err
}

func DeCompress(dest string, src string) error {

	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	newfile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer newfile.Close()

	zr, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	filestat, err := file.Stat()
	if err != nil {
		return err
	}

	zr.Name = filestat.Name()
	zr.ModTime = filestat.ModTime()
	if _, err := io.Copy(newfile, zr); err != nil {
		return err
	}
	if err := zr.Close(); err != nil {
		return err
	}
	return nil
}
