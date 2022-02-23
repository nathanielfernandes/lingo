package lingo

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"

	gogitignore "github.com/sabhiram/go-gitignore"
)

type File struct {
	Language string
	Path     string
}

func GetIdentity(fn string) (string, bool) {
	for name, id := range LANGUAGE_IDENTIFIERS {
		if id.Matches(fn) {
			return name, true
		}
	}
	return "", false
}

func GetLanguages(start string, gitignore bool) Languages {
	langs := Languages{}

	for _, f := range GetFilePaths(start, gitignore) {
		lang, ok := langs[f.Language]
		if ok {
			lang.Files = append(lang.Files, f.Path)
		} else {
			langs[f.Language] = &Language{GetAlias(f.Language), []string{f.Path}, 0, GetColor(f.Language)}
		}
	}

	return langs
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func GetFilePaths(start string, gitignore bool) []File {
	paths := []File{}
	ignorePath := ""

	err := filepath.Walk(start,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.HasPrefix(path, ".git") || info.IsDir() {
				if strings.HasSuffix(path, ".gitignore") {
					ignorePath = path
				}

				return nil
			}

			paths = append(paths, File{info.Name(), path})
			return nil
		})

	if err != nil {
		log.Println(err)
	}

	filePaths := []File{}
	if gitignore && ignorePath != "" {
		lines, err := ReadLines(ignorePath)
		if err != nil {
			log.Println(err)
		} else {
			object := gogitignore.CompileIgnoreLines(lines...)

			for _, f := range paths {
				if !object.MatchesPath(f.Path) {
					lang, ok := GetIdentity(f.Language)
					if !ok {
						continue
					}
					filePaths = append(filePaths, File{lang, f.Path})
				}
			}
		}
	} else {
		for _, f := range paths {
			lang, ok := GetIdentity(f.Language)
			if !ok {
				continue
			}
			filePaths = append(filePaths, File{lang, f.Path})
		}
	}

	return filePaths
}

func GetFilePath1s(start string) []File {
	FilePaths := []File{}

	err := filepath.Walk(start,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.HasPrefix(path, ".git") || info.IsDir() {
				return nil
			}

			lang, ok := GetIdentity(info.Name())
			if !ok {
				return nil
			}

			FilePaths = append(FilePaths, File{lang, path})
			return nil
		})

	if err != nil {
		log.Println(err)
	}

	return FilePaths
}
