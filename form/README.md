<p align="center">
  <img src="https://cdn.aahframework.org/assets/img/aah-logo-64x64.png" />
  <h2 align="center">Example - Form Submission</h2>
</p>

This example demonstrates how form submission happens in aah framework. Form submission includes auto parse and bind (nested struct bind and input sanitization to prevent XSS attacks). Take a moment to learn about the clean and robust implementation of [auto parse and bind](https://docs.aahframework.org/request-parameters-auto-bind.html) capabilities.

### Get aah examples

```bash
git clone https://github.com/go-aah/examples.git $GOPATH/src/aahframework.org/examples
```

### Run this example

```bash
aah r -i aahframework.org/examples/form
```

### Visit this URL

  * http://localhost:8080

Home page includes the details of implemented functionality and button to access User Profile - Form Submission. Fill out the form fields and click submit. As a response, the submitted field values will be displayed on the page.
