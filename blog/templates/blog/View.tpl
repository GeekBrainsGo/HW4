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
            {{template "Item" .}}
        </div>
    </body>
    {{template "Resources"}}
    {{template "Scripts"}}
</html>
{{end}}

{{define "Item"}}
<div class="uk-card uk-card-body">
	<div class="uk-card uk-card-default uk-card-body">
        <h4 post-id="{{.Data.ID}}">{{.Data.Title}}</h4>
        <span>{{.Data.Short}}</span>
        <div class="uk-align-right">
			{{.Data.Body}}
        </div>
	</div>
</div>
{{end}}