package resource

import (
	"fmt"
	"os"
)

const (
	MAPVAR = "ResMap"
)

func Convert(srcDir string, dstFilePath string, packageName string) error {
	var err error
	var file *os.File

	if file, err = os.OpenFile(dstFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644); err != nil {
		return err
	}
	defer file.Close()

	if _, err = fmt.Fprintf(file, "package %s\n\n", packageName); err != nil {
		return err
	}

	if _, err = fmt.Fprintf(file, "var %s = map[string] string {}\n\n", MAPVAR); err != nil {
		return err
	}

	if err = ConvertDir(srcDir, "", file); err != nil {
		return err
	}

	return nil
}
