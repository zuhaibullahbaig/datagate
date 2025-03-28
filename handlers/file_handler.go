package handlers

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
)

func ListFiles(directory string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Println("Error reading directory:", err)
		return nil, err
	}
	return files, nil
}

func IsDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func GetFilePath(baseDir, name string) string {
	return filepath.Join(baseDir, name)
}

func GetFileExtension(filename string) string {
	return filepath.Ext(filename)
}

func FormatFileSize(size int64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)

	switch {
	case size >= GB:
		return fmt.Sprintf("%.2f GB", float64(size)/float64(GB))
	case size >= MB:
		return fmt.Sprintf("%.2f MB", float64(size)/float64(MB))
	case size >= KB:
		return fmt.Sprintf("%.2f KB", float64(size)/float64(KB))
	default:
		return fmt.Sprintf("%d B", size)
	}
}

func ConstructPath(currentDir, fileName string, isDir bool, token string, authEnabled bool) string {
	var relativePath string
	if currentDir == "" {
		relativePath = fileName
	} else {
		relativePath = filepath.Join(currentDir, fileName)
	}

	relativePath = filepath.ToSlash(relativePath)

	var path string
	if isDir {
		path = fmt.Sprintf("/open?dir=%s", url.PathEscape(relativePath))
	} else {
		path = fmt.Sprintf("/download?file=%s", url.PathEscape(relativePath))
	}

	if authEnabled {
		path += "&token=" + token
	}

	return path
}
