package filepath

import "strings"

func Ext(path string, defaultExt string) string {
	l := len(path)
	if path[l:l] == "/" {
		return defaultExt
	}
	if strings.ContainsRune(path, '.') {
		return path[strings.LastIndexByte(path, '.')+1 : l]
	}
	return defaultExt
}

func Path(path string) string {
	if strings.ContainsRune(path, '/') {
		return path[:strings.LastIndexByte(path, '/')]
	}
	return ""
}
