<p align="center">
  <img src="https://cdn.aahframework.org/assets/img/aah-logo-64x64.png" />
  <h2 align="center">Example - i18n URL Query Parameter</h2>
</p>

This example demonstrates i18n - Internationalization with aah framework using **URL Query Parameter**. It is by applying aah configuration and convention. Zero coding effort on this aspect!

Application has URL pattern as 

```
/?lang={language}
```

Learn more on how i18n works in aah - https://docs.aahframework.org/i18n.html.

## How to get the aah examples?

```bash
git clone https://github.com/go-aah/examples.git $GOPATH/src/aahframework.org/examples
```

## How to run this example?

```bash
aah r -i aahframework.org/examples/i18n-url-query-param
```

### Now visit this URL

Go to home page click on the language links to see i18n in action.

  * Home page : http://localhost:8080

Localized URLs:

- http://localhost:8080/?lang=en
- http://localhost:8080/?lang=zh-CN
- http://localhost:8080/?lang=ja
- http://localhost:8080/?lang=es-ES
- http://localhost:8080/?lang=es-MX
- http://localhost:8080/?lang=de-DE

