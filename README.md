# 目的
- REST APIとVIEWを分け、フレキシブルかつ柔軟なウェブサイトにする。
- 初期設定を簡素化する。
- html/templateのBlockでいつも詰まるので、ボイルプレートに

# Usage
- ホットリロード（Hot Reload）
``` command
$ realize run
```

- html/template（block）
React/Angular/Vueなどでも活用されているテンプレートとページファイルの分割を実現。./\_views/pagesにファイルを増やし、contorollersで呼び出せばOK。  
pagesに追加するファイルは、前置詞（primary-/secondary-/admin-）でレイアウトテンプレートを選択でき、その他前置詞の場合は./\_views/\_layouts/default.gohtmlが呼ばれます。
単一ファイル内でhtml/script/styleを分けました。同一ファイルに読み込まれ、styleはhead内、scriptはfooter下に配置しています。

``` html
<!-- ここはレイアウトテンプレート -->
<!doctype html>
<html lang="ja-JP">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Primary</title>
    <link href="" rel="shortcut icon">
    <!-- # Common styles and more -->
    <link href="" rel="stylesheet" type="text/css" />

    <!-- # og -->
    <meta property="og:type" content="website">
    <meta property="og:image" content="">
    <meta property="og:site_name" content="">
    <meta property="og:title" content="">
    <meta property="og:url" content="">

    <meta name="twitter:card" content="">
    <meta name="twitter:image" content="">
    <meta name="twitter:site" content="">
    <meta name="twitter:creator" content="">
    <meta name="twitter:title" content="">
    <meta name="twitter:url" content="">

    <!-- # Load individual styles -->
    {{block "styles" .}}Styles{{end}}
  </head>
<body>
  <!-- # Primary header -->

  {{block "main" .}}Main{{end}}

  <footer>
    footer
  </footer>

  {{block "scripts" .}}Scripts{{end}}
</body>
</html>

```

``` html
<!-- ここはpage -->
<!-- 上記レイアウトテンプレートにバシッと挿入されます -->
{{define "main"}}
<section class="section">
  <div class="container">
    This is format Template.
  </div>
</section>
{{end}}


{{define "scripts"}}
<script>

</script>
{{end}}


{{define "styles"}}
<style scoped>

</style>
{{end}}
```

- Components にて部品開発
./\_views/\_componets を配置。内部で再利用可能性が高いコンポーネントを開発し、tempalteで呼び出してください。

``` html
{{define "component"}}
// something
{{end}}
```

``` html
{{template "component"}}
```

- 静的ディレクトリ /\_statics  
静的ディレクトリには、現状html/templateから呼び出すcss/js/imagesフォルダを配置しました。htmlからsrc="/images/thisfile.jpg"で呼び出せます。
ちなみに、ReactやAnguler・Vueをここにいれて開発することが可能です。GolangにAPI専業させて、ViewはjavaScriptでヌルヌル動かしたいなどは、このディレクトリか./app内で

## Bio
K-Terashima
- Twitter: [@6rules](http://twitter.com/6rules)
