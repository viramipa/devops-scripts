package devops_scripts

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GenerateRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err!= nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func EnsureDir(path string) error {
	return os.MkdirAll(path, 0755)
}

func GetDir(path string) string {
	return filepath.Dir(path)
}

func GetFileExtension(path string) string {
	return filepath.Ext(path)
}

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return!os.IsNotExist(err)
}

func GetFileSize(path string) (int64, error) {
	info, err := os.Stat(path)
	if err!= nil {
		return 0, err
	}
	return info.Size(), nil
}

func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err!= nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err!= nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

func ReadFileContents(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err!= nil {
		return "", err
	}
	return string(data), nil
}

func WriteFileContents(path string, contents string) error {
	return os.WriteFile(path, []byte(contents), 0644)
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func LogError(err error) {
	log.Println(err.Error())
}

func LogErrorf(format string, a...interface{}) {
	log.Printf(format, a...)
}

func LogInfo(message string) {
	log.Println(message)
}

func LogInfof(format string, a...interface{}) {
	log.Printf(format, a...)
}

func IsEmptyString(str string) bool {
	return str == ""
}

func TrimString(str string) string {
	return strings.TrimSpace(str)
}

func GetEnvVariable(key string) string {
	return os.Getenv(key)
}

func SetEnvVariable(key, value string) {
	os.Setenv(key, value)
}

func GetHomeDir() string {
	return os.Getenv("HOME")
}

func GetRandomPort() (int, error) {
	l, err := net.Listen("tcp", ":0")
	if err!= nil {
		return 0, err
	}
	defer l.Close()

	return l.Addr().(*net.TCPAddr).Port, nil
}

func GetHostname() string {
	return os.Getenv("HOSTNAME")
}

func GetProcessID() int {
	return os.Getpid()
}

func GetExecutablePath() string {
	return os.Args[0]
}