package bindata

import (
	"github.com/Unknwon/macaron"
	"github.com/elazarl/go-bindata-assetfs"
)

type (
	templateFileSystem struct {
		files []macaron.TemplateFile
	}

	templateFile struct {
		name string
		data []byte
		ext  string
	}

	Options struct {
		// Asset should return content of file in path if exists
		Asset func(path string) ([]byte, error)
		// AssetDir should return list of files in the path
		AssetDir func(path string) ([]string, error)
		// AssetNames should return list of all asset names
		AssetNames func() []string
		// Prefix would be prepended to http requests
		Prefix string
	}
)

func Static(opt Options) *assetfs.AssetFS {
	fs := &assetfs.AssetFS{
		Asset:    opt.Asset,
		AssetDir: opt.AssetDir,
		Prefix:   opt.Prefix,
	}

	return fs
}

func (templates templateFileSystem) ListFiles() []macaron.TemplateFile {
	return templates.files
}

func (f *templateFile) Name() string {
	return f.name
}

func (f *templateFile) Data() []byte {
	return f.data
}

func (f *templateFile) Ext() string {
	return f.ext
}

func Templates(opt Options) templateFileSystem {
	fs := templateFileSystem{}
	fs.files = make([]macaron.TemplateFile, 0, 10)

	list := opt.AssetNames()

	for _, key := range list {
		ext := macaron.GetExt(key)

		data, err := opt.Asset(key)

		if err != nil {
			continue
		}

		name := (key[0 : len(key)-len(ext)])

		fs.files = append(fs.files, &templateFile{name, data, ext})
	}

	return fs
}
