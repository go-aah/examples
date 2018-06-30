<p align="center">
  <img src="https://cdn.aahframework.org/assets/img/aah-logo-64x64.png" />
  <h2 align="center">Example - Form Submission</h2>
</p>

This example demonstrates Form Submission with aah framework. That includes Auto parse and bind (nested struct bind and input sanitization to prevent XSS attacks). Take a moment to learn about clean, robust implementation of [auto parse and bind](https://docs.aahframework.org/request-parameters-auto-bind.html) capabilities.

## How to get the aah examples?

```bash
git clone https://github.com/go-aah/examples.git $GOPATH/src/aahframework.org/examples
```

## How to run this example?

```bash
aah r -i aahframework.org/examples/form
```

### Now visit this URL

  * http://localhost:8080

Home page includes details about the implemented functionality and button to access User Profile - Form Submission. Fill out the form fields and submit. As a response, submitted field values displayed on the page.