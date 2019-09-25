{{define "page"}}
<!DOCTYPE html>
<html>
    <head>
        {{template "Head"}}
        <title>{{.Title}}</title>
        
    </head>
    <body>
        <div class="uk-card uk-card-default uk-card-body">
            <h3>{{.Title}}</h3>
            <span>All: {{len .Data}}</span>
            {{template "Items" .}}
        </div>
    </body>
    {{template "Resources"}}
    {{template "Scripts"}}
</html>
{{end}}

{{define "Items"}}
<div class="uk-card uk-card-body">
    <ul class="uk-list">
    {{range $key, $val := .Data}}
        <li>
            <div class="uk-card uk-card-default uk-card-body">
                <a href="/view/{{.ID}}"><h4 post-id="{{.ID}}">{{.Title}}</h4></a>
                <span>{{.Short}}</span>
                <div class="uk-align-right">
					{{.Body}}
                </div>
            </div>
        </li>
    {{end}}
    </ul>
</div>
{{end}}