package util

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func EncryptAndCompressDir(src string, buf io.Writer, key string) error {
	// aes > tar > gzip > buf
	ew, err := EncryptedWriter(key, buf)
	if err != nil {
		return err
	}
	zr := gzip.NewWriter(ew)
	tw := tar.NewWriter(zr)

	// is file a folder?
	fi, err := os.Stat(src)
	if err != nil {
		return err
	}
	mode := fi.Mode()
	if !mode.IsDir() {
		return fmt.Errorf("error: file type not supported")
	}
	err = filepath.Walk(src, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if file == src {
			return nil
		}

		var link string
		if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
			if link, err = os.Readlink(file); err != nil {
				return err
			}
		}

		header, err := tar.FileInfoHeader(fi, link)
		if err != nil {
			return err
		}

		header.Name = filepath.Join(filepath.Base(src), strings.TrimPrefix(file, src))
		if err = tw.WriteHeader(header); err != nil {
			return err
		}

		if !fi.Mode().IsRegular() { //nothing more to do for non-regular
			return nil
		}

		fh, err := os.Open(file)
		if err != nil {
			return err
		}
		defer fh.Close()

		if _, err = io.Copy(tw, fh); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	// produce tar
	if err := tw.Close(); err != nil {
		return err
	}
	// produce gzip
	if err := zr.Close(); err != nil {
		return err
	}
	//
	return nil
}
