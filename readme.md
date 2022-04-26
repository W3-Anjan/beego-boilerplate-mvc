## Project Run 
put this url in browser: http://localhost:8080/index otherwise builtin error will occur

### Go lang BeeGo framework basic and useful things 

### Key reference 
1. https://github.com/vv1133/vvblog4
2. https://www.educative.io/edpresso/what-is-html-template-in-golang

- ``{{range .Todos}} {{.}} {{end}}	Loops over all “Todos” and renders each using {{.}}``
- ``{{block "content" .}} {{end}} Defines a block with the name “content”``


### Controller define Error
1. https://github.com/beego/beedoc/blob/master/en-US/mvc/controller/errors.md#controller-define-error

### ne eq
text/template doesn’t provide operators, but it provides functions such as and, or, not, eq or ne.
The html/template package provides us with a few classes to help do comparison. These are
- ``eq`` - Returns the boolean truth of ``arg1 == arg2``
- ``ne`` - Returns the boolean truth of ``arg1 != arg2``
- ``lt`` - Returns the boolean truth of ``arg1 < arg2``
- ``le`` - Returns the boolean truth of ``arg1 <= arg2``
- ``gt`` - Returns the boolean truth of ``arg1 > arg2``
- ``ge`` - Returns the boolean truth of ``arg1 >= arg2``

### Filter
Beego supports custom filter middlewares. E.g.: user authentication and force redirection

### Flash
Flash messages are temporary messages between two logic blocks. All flash messages will be cleared after the very next logic block. They are normally used to send notes and error messages.