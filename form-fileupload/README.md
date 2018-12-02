<p align="center">
  <img src="https://cdn.aahframework.org/assets/img/aah-logo-64x64.png" />
  <h2 align="center">Example - Form File Upload</h2>
</p>

This example illustrates aah form file upload. It is easy to handle file uploads in aah. It works by calling (method) `Req.SaveFile` with form field name and destination path. The method returns file size on successful save.

```go
fileSize, err := Req.SaveFile("userProfileImage", "/Users/jeeva/user-profile-uploads")
```

### Get aah examples

```bash
git clone https://github.com/go-aah/examples.git aah-examples
```

### Run this example

```bash
cd aah-examples/form-fileupload
aah run
```

### Visit this URL

  * http://localhost:8080

Home page includes the details of implemented functionality and has a button to choose file and upload. It is self explanatory.
