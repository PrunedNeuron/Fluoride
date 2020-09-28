package redoc

import (
	"bytes"
	"html/template"
	"net/http"
	"path"

	"github.com/PrunedNeuron/Fluoride/config"
)

// New creates a middleware to serve a documentation site for a swagger spec.
// This allows for altering the spec before starting the http listener.
func New(next http.Handler) http.Handler {

	options := config.GetConfig().Documentation

	pth := path.Join(options.BasePath, options.Path)
	tmpl := template.Must(template.New("redoc").Parse(redocTemplate))

	buf := bytes.NewBuffer(nil)
	_ = tmpl.Execute(buf, options)
	b := buf.Bytes()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == pth {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusOK)

			_, _ = w.Write(b)
			return
		}
		filePath := options.BasePath + options.Path + options.SpecURL

		/* if _, err := validate.OpenAPI(filePath); err != nil {
			zap.S().Errorf("Failed to validate Open API specification, error: " + err.Error())
		} */

		// serve spec if path = docs/specPath.ext
		if r.URL.Path == pth+options.SpecURL {
			http.ServeFile(w, r, "."+filePath)
			return
		}
		next.ServeHTTP(w, r)
	})
}

const (
	redocTemplate = `<!DOCTYPE html>
<html>
  <head>
    <title>{{ .Title }}</title>
    <!-- needed for adaptive design -->
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!--
    ReDoc doesn't change outer page styles
    -->
    <style>
      body {
        margin: 0;
        padding: 0;
      }
    </style>
  </head>
  <body>
    <redoc spec-url='{{ .Path }}{{ .SpecURL }}'></redoc>
    <script src="{{ .RedocURL }}"></script>
  </body>
</html>
`
)
