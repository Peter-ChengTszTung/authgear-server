package web

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"regexp"

	"github.com/authgear/authgear-server/pkg/util/resource"
	"github.com/authgear/authgear-server/pkg/util/template"
)

type languageImage struct {
	languageTag string
	data        []byte
}

func (i languageImage) GetLanguageTag() string {
	return i.languageTag
}

var imageExtensions = map[string]string{
	"image/png":  ".png",
	"image/jpeg": ".jpeg",
	"image/gif":  ".gif",
}

var imageRegex = regexp.MustCompile(`^(.+)_([a-zA-Z0-9-]+)\.(png|jpeg|gif)$`)

const argResolvedLanguageTag = "resolved_language_tag"

type imageAsset struct {
	Name string
}

func (a imageAsset) ReadResource(fs resource.Fs) ([]resource.LayerFile, error) {
	dir := path.Dir(a.Name)
	fileNames, err := resource.ReadDirNames(fs, dir)
	if os.IsNotExist(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	langFiles := make(map[string]resource.LayerFile)
	for _, fileName := range fileNames {
		p := path.Join(dir, fileName)
		matches := imageRegex.FindStringSubmatch(p)
		if len(matches) != 4 {
			continue
		}
		name := matches[1]
		languageTag := matches[2]
		if name != a.Name {
			continue
		}
		if f, ok := langFiles[languageTag]; ok {
			return nil, fmt.Errorf("duplicated image files: %s, %s", f.Path, p)
		}

		data, err := resource.ReadFile(fs, p)
		if os.IsNotExist(err) {
			continue
		} else if err != nil {
			return nil, err
		}

		langFiles[languageTag] = resource.LayerFile{Path: p, Data: data}
	}

	var files []resource.LayerFile
	for _, file := range langFiles {
		files = append(files, file)
	}
	return files, nil
}

func (a imageAsset) MatchResource(path string) bool {
	matches := imageRegex.FindStringSubmatch(path)
	if len(matches) != 4 {
		return false
	}
	return matches[1] == a.Name
}

func (a imageAsset) Merge(layers []resource.LayerFile, args map[string]interface{}) (*resource.MergedFile, error) {
	preferredLanguageTags, _ := args[ResourceArgPreferredLanguageTag].([]string)
	defaultLanguageTag, _ := args[ResourceArgDefaultLanguageTag].(string)
	// If user requested static asset at a specific path, always use the
	// corresponding language in path
	if p, ok := args[ResourceArgRequestedPath].(string); ok {
		match := imageRegex.FindStringSubmatch(p)
		if len(match) == 4 {
			languageTag := match[2]
			preferredLanguageTags = []string{languageTag}
		}
	}

	images := make(map[string]template.LanguageItem)
	for _, file := range layers {
		languageTag := imageRegex.FindStringSubmatch(file.Path)[2]
		images[languageTag] = languageImage{
			languageTag: languageTag,
			data:        file.Data,
		}
	}

	var items []template.LanguageItem
	for _, i := range images {
		items = append(items, i)
	}

	matched, err := template.MatchLanguage(preferredLanguageTags, defaultLanguageTag, items)
	if errors.Is(err, template.ErrNoLanguageMatch) {
		if len(items) > 0 {
			// Use first item in case of no match, to ensure resolution always succeed
			matched = items[0]
		} else {
			// If no configured translation, fail the resolution process
			return nil, resource.ErrResourceNotFound
		}
	} else if err != nil {
		return nil, err
	}

	tagger := matched.(languageImage)
	return &resource.MergedFile{
		Args: map[string]interface{}{
			argResolvedLanguageTag: tagger.languageTag,
		},
		Data: tagger.data,
	}, nil
}

func (a imageAsset) Parse(merged *resource.MergedFile) (interface{}, error) {
	mimeType := http.DetectContentType(merged.Data)
	ext, ok := imageExtensions[mimeType]
	if !ok {
		return nil, fmt.Errorf("invalid image format: %s", mimeType)
	}

	var path string
	if langTag, ok := merged.Args[argResolvedLanguageTag]; ok {
		path = fmt.Sprintf("%s_%s%s", a.Name, langTag, ext)
	} else {
		path = a.Name + ext
	}

	return &StaticAsset{
		Path: path,
		Data: merged.Data,
	}, nil
}
