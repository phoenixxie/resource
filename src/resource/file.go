package resource

import (
	"bytes"
	"compress/flate"
	"encoding/ascii85"
	"fmt"
	"io"
	"os"
	"strconv"
)

func readFile(filePath string) (string, error) {
	var buf []byte
	var err error
	var n int
	var writer *flate.Writer
	var encoder io.WriteCloser
	var file *os.File
	var data *bytes.Buffer = new(bytes.Buffer)

	if file, err = os.Open(filePath); err != nil {
		return "", err
	}
	defer file.Close()

	encoder = ascii85.NewEncoder(data)
	defer encoder.Close()

	writer, err = flate.NewWriter(encoder, flate.BestCompression)
	defer writer.Close()

	buf = make([]byte, 8192)
	err = nil
	for {
		if n, err = file.Read(buf); err != nil {
			if err == io.EOF {
				break
			} else {
				return "", err
			}
		}

		if n == 8192 {
			_, err = writer.Write(buf)
		} else {
			_, err = writer.Write(buf[:n])
		}
		if err != nil {
			return "", err
		}
	}
	if err = writer.Flush(); err != nil {
		return "", err
	}

	return strconv.Quote(string(data.Bytes())), nil
}

func ConvertFile(filePath string, key string, writer io.Writer) error {
	var err error = nil
	var data string

	if data, err = readFile(filePath); err != nil {
		return err
	}

	_, err = fmt.Fprintf(writer, "%s[\"%s\"] = %s\n", MAPVAR, key, data)

	return err
}
