{{define "nav"}}
	<div class="navbar navbar-default navbar-fixed-top">
  		<div class="container">
  			<a href="/" class="navbar-brand">my blog</a>
  			<ul class="nav navbar-nav">
				<li {{if .isHome}}class="active"{{end}}><a href="/">Home</a></li>
				<li {{if .isCategory}}class="active"{{end}}><a href="/category">category</a></li>
				<li {{if .isTopic}}class="active"{{end}}><a href="/topic">topic</a></li>
			</ul>
			<ul class="nav navbar-nav pull-right">
				{{if .isLogin}}
				<li><a href="/login?exit=true">logout</a></li>
				{{else}}
				<li><a href="/login">login</a></li>
				{{end}}
			</ul>
	    </div>
	</div>
{{end}}