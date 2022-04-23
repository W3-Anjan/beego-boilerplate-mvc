<!DOCTYPE html>
<html>
{{ block "head" . }}{{ end }}
{{ block "css" . }}{{ end }}
<body>
    <div class="container">
        {{ block "content" . }}{{ end }}
    </div>
    <script type="text/javascript" src="http://code.jquery.com/jquery-2.0.3.min.js"></script>
    <script src="http://netdna.bootstrapcdn.com/bootstrap/3.0.3/js/bootstrap.min.js"></script>
    <script src="/static/js/reload.min.js"></script>
     {{ block "js" . }}{{ end }}
</body>
</html>