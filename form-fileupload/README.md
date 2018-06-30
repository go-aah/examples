<p align="center">
  <img src="https://cdn.aahframework.org/assets/img/aah-logo-64x64.png" />
  <h2 align="center">Example - Form File Upload</h2>
</p>

This example demonstrates Form File Upload with aah framework. It is easy to handle File uploads in aah, typically calling `Req.SaveFile` with form field name and destination path. Method returns file size on successful save.

```go
fileSize, err := Req.SaveFile("userProfileImage", "/Users/jeeva/user-profile-uploads")
```

## How to get the aah examples?

```bash
git clone https://github.com/go-aah/examples.git $GOPATH/src/aahframework.org/examples
```

## How to run this example?

```bash
aah r -i aahframework.org/examples/form-fileupload
```

### Now visit this URL

  * http://localhost:8080

Home page includes details about the implemented functionality and button to choose file and upload. It is self explanatory.