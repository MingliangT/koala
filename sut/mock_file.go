package sut

import (
	"crypto/sha1"
	"encoding/base32"
	"io/ioutil"
	"github.com/v2pro/koala/countlog"
	"os"
	"github.com/v2pro/koala/envarg"
)

func init() {
	if envarg.IsReplaying() {
		os.Mkdir("/tmp/koala-mocked-files", 0777)
	}
}

func mockFile(content []byte) string {
	mockedFile := "/tmp/koala-mocked-files/" + hash(content)
	if _, err := os.Stat(mockedFile); err == nil {
		return mockedFile
	}
	err := ioutil.WriteFile(mockedFile, content, 0666)
	if err != nil {
		countlog.Error("event!sut.failed to write mock file",
			"mockedFile", mockedFile, "err", err)
		return ""
	}
	return mockedFile
}

func hash(content []byte) string {
	h := sha1.New()
	h.Write(content)
	return "g" + base32.StdEncoding.EncodeToString(h.Sum(nil))
}
