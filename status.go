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
<style>
table {
  width: 100%;
  text-align: left;
  border-collapse: collapse;
}
table td, table th {
  border: 1px solid black;
  padding: 2px 2px;
}
table tr:nth-child(even) {
  background: #EEEEEE;
}
table thead {
  border-bottom: 4px solid #333333;
}
table thead th {
  font-weight: bold;
}
ul {
  padding-left: 0px;
}
</style>
</head>
<body>
<p>
<h1>Sitest</h1>
<table>
<tr><th>Site</th><th>Config</th><th>Last result</th></tr>
{{ range $key, $value := .Sites }}
   <tr><td>{{ $key }}</td><td>Interval: {{ $value.Config.Interval }}</td><td>{{LastResultFormated .}}</td></tr>
{{ end }}
</table>
</p>
<p class="footer">
<ul>
<li>Configuration file: {{.ConfigFile}}</li>
<li>Start time {{ .StartTime }}</li>
<li>Detailed metrics: <a href="/metrics">/metrics</a></li>
</ul>
</p>
</body>
</html>`
	var funcs = template.FuncMap{"LastResultFormated": func(site *Site) Result { return site.GetLastResult() }}

	tmpl, err := template.New("status-page").Funcs(funcs).Parse(tmplTxt)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, sitest)
	if err != nil {
		log.Fatal(err)
	}
}
