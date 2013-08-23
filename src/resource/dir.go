package resource

import (
	"io"
	"io/ioutil"
	"os"
	"path"
)

func ConvertDir(dir string, keyRoot string, writer io.Writer) error {
	var err error = nil
	var entries []os.FileInfo

	if entries, err = ioutil.ReadDir(dir); err != nil {
		return err
	}

	for _, entry := range entries {
		fullPath := path.Join(dir, entry.Name())
		keyBase := keyRoot + "/" + entry.Name()

		if entry.IsDir() {
			err = ConvertDir(fullPath, keyBase, writer)
		} else {
			err = ConvertFile(fullPath, keyBase, writer)
		}
		if err != nil {
			return err
		}
	}

	return nil
}
