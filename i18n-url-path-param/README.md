<p align="center">
  <img src="https://cdn.aahframework.org/assets/img/aah-logo-64x64.png" />
  <h2 align="center">Example - i18n URL Path Parameter</h2>
</p>

This example demonstrates i18n - Internationalization with aah framework using **URL Path Parameter**. By applying aah configuration and convention, zero coding effort can be achieved on this aspect.

Application has below route URL pattern -

```
/page/:language/index.html
```

Learn more on how i18n works in aah - https://docs.aahframework.org/i18n.html.

### Get aah examples

```bash
git clone https://github.com/go-aah/examples.git $GOPATH/src/aahframework.org/examples
```

### Run this example

```bash
aah r -i aahframework.org/examples/i18n-url-path-param
```

### Visit this URL

Go to home page, click on language links to see i18n in action.

  * Home page : http://localhost:8080

Localized URLs:

- http://localhost:8080/page/en/index.html
- http://localhost:8080/page/zh-CN/index.html
- http://localhost:8080/page/ja/index.html
- http://localhost:8080/page/es-ES/index.html
- http://localhost:8080/page/es-MX/index.html
- http://localhost:8080/page/de-DE/index.html 

