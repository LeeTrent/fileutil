package fileutil

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func WriteFileToDisk(srcFile io.Reader, dirName string, fileName string) error {

	// read source file
	bs, err := ioutil.ReadAll(srcFile)
	if err != nil {
		return err
	}

	// create destination file
	dstFile, err := os.Create(filepath.Join(dirName, fileName))
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// write destination file to disk
	_, err = dstFile.Write(bs)
	if err != nil {
		return err
	}

	return nil
}

func ReadFileFromDisk(dirName string, fileName string) (records [][]string, err error) {

	data, err := os.Open(filepath.Join(dirName, fileName))
	if err != nil {
		return nil, err
	}
	defer data.Close()

	reader := csv.NewReader(data)
	recs, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return recs, nil
}
