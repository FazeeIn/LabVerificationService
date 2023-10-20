package model

import (
	"archive/tar"
	"bytes"
)

func newFile(name string, data []byte) (bytes.Buffer, error) {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	err := tw.WriteHeader(&tar.Header{
		Name: name,             // filename
		Mode: 0644,             // permissions
		Size: int64(len(data)), // filesize
	})
	if err != nil {
		return buf, err
	}
	tw.Write(data)
	tw.Close()
	return buf, nil
}
