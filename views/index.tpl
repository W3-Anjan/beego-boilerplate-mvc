{{ template "layout_blog.tpl" . }}
{{ define "css" }}
		<link rel="stylesheet" href="/static/css/current.css">
{{ end}}


{{ define "content" }}
		 <header>
            <h1 class="logo">Welcome to Beego</h1>
            <div class="description">
              Beego is a simple & powerful Go web framework which is inspired by tornado and sinatra.
            </div>
            <div>Request URL:{{ $.RequestUrl }}</div>
            <div> Beego built-in template func date: {{date .t "Y-m-d H:i:s"}} </div>
            <div> Beego built-in template func map_get: {{ map_get .m "a"}} & {{ map_get .m 1 "c" }}</div>
            <div> Beego Custom template func LoadTimes: {{.t | LoadTimes }}</div>
          </header>
          <footer>
             <div class="author">
                  Official website:
                  <a href="http://{{.Website}}">{{.Website}}</a> /
                  Contact me:
                  <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
              </div>
          </footer>
{{ end }}

{{ define "js" }}
	<script src="/static/js/current.js"></script>
{{ end}}