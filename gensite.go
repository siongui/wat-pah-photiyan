package gensite

import (
	"github.com/siongui/gotm"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func indexHtml(tmpldir, outputdir string, data interface{}) (err error) {
	tm := gotm.NewTemplateManager("")
	err = tm.ParseDirectory(tmpldir)
	if err != nil {
		return
	}

	fo, err := os.Create(path.Join(outputdir, "index.html"))
	if err != nil {
		return err
	}
	defer fo.Close()

	return tm.ExecuteTemplate(fo, "index.html", data)
}

func IsHtml(info os.FileInfo) bool {
	if info.IsDir() {
		return false
	}
	return strings.HasSuffix(info.Name(), ".html")
}

func ParseContentDir(dir string) {
	// walk all files in directory
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if IsHtml(info) {
			ParseHtmlContent(path)
		}
		return nil
	})
}
