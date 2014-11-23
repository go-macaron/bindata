# bindata 

Package bindata is a helper module that allows to use in-memory static and template files for Macaron via go-bindata (https://github.com/jteeuwen/go-bindata)

Using go-bindata convert your template and public directories into individual packages. 
Import the packages and use them like the example below.

## Example
```go
import "path/to/bindata/public"
import "path/to/bindata/templates"

m.Use(macaron.Static(
  path.Join(setting.StaticRootPath, "public"),
  macaron.StaticOptions{
    SkipLogging: false,
    FileSystem: bindata.Static(bindata.Options{
      Asset:      public.Asset,
      AssetDir:   public.AssetDir,
      AssetNames: public.AssetNames,
      Prefix:     "",
    }),
  },
))

m.Use(macaron.Renderer(macaron.RenderOptions{
  Funcs:      []template.FuncMap{base.TemplateFuncs},
  IndentJSON: macaron.Env != macaron.PROD,
  TemplateFileSystem: bindata.Templates(bindata.Options{
    Asset:      templates.Asset,
    AssetDir:   templates.AssetDir,
    AssetNames: templates.AssetNames,
    Prefix:     "",
  }),
}))
```

# MIT LICENSE

The MIT License (MIT) - [LICENSE](https://github.com/mephux/dnas/blob/master/LICENSE)
