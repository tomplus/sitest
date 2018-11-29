package main

import (
	"html/template"
	"log"
	"net/http"
)

func (sitest Sitest) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	tmplTxt := `<html>
<head>
<title>Sitest</title>
</head>
<body>
<p>
<h1>Defined sites:</h1>
<ul>
{{ range $key, $value := .Sites }}
   <li><strong>{{ $key }}</strong>: {{ $value }}</li>
{{ end }}
</ul>
configuration from file {{.ConfigFile}}</pre>
</p>
<p>
Metrics <a href="/metrics">/metrics</a>
</p>
</body>
</html>`

	tmpl, err := template.New("status-page").Parse(tmplTxt)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, sitest)
	if err != nil {
		log.Fatal(err)
	}
}
