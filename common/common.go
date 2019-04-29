package common

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	// ServersRoot is the root directory where all of the servers are stored.
	ServersRoot = "servers/"
)

var (
	// ErrServernameAlreadyInUse is thrown when there is already a server with that name
	ErrServernameAlreadyInUse = errors.New("that server name is already in use")
)

// CreateDirIfDoesNotExist creates a directory if it does not already exist.
func CreateDirIfDoesNotExist(dir string) error {
	dir = filepath.FromSlash(dir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// NewServerPath returns the path to a server given its name.
func NewServerPath(name string) (string, error) {
	// Create a path
	rawPath := ServersRoot + name
	abs, err := filepath.Abs(rawPath)
	if err != nil {
		return "", err
	}

	return abs, nil

}

// Unzip decompresses a zip archive, moving all files and folders
// within the zip file (src) to an output directory (dest).
func Unzip(src string, dest string) ([]string, error) {
	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}
