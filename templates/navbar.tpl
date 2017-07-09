{{ define "navbar" }}
<nav>
  <div class="nav-wrapper">
    <a href="/" class="brand-logo left">Bayo Bilac</a>
    <ul class="right">
      <li {{ if eq .site "table" }} class="active"{{ end }}><a href="/">Table</a></li>
      <li {{ if eq .site "elo" }} class="active"{{ end }}><a href="/elo">Elo</a></li>
      <li {{ if eq .site "draw" }} class="active"{{ end }}><a href="/draw">Draw</a></li>
    </ul>
  </div>
</nav>
{{ end }}
