<p align="center">
  <img src="https://cdn.aahframework.org/assets/img/aah-logo-64x64.png" />
  <h2 align="center">Example - Form based Auth</h2>
</p>

This example demonstrates Form based Auth with aah framework. That includes Authentication, Route Authorization via routes config, Access permissions on View files, Session access on View files.

Learn more about [Security design](https://docs.aahframework.org/security-design.html), [Authentication](https://docs.aahframework.org/authentication.html), [Authorization](https://docs.aahframework.org/authorization.html) and [Form Auth scheme](https://docs.aahframework.org/auth-schemes/form.html).

## How to get the aah examples?

```bash
git clone https://github.com/go-aah/examples.git $GOPATH/src/aahframework.org/examples
```

## How to run this example?

```bash
aah r -i aahframework.org/examples/form-based-auth
```

### Now visit this URL

  * http://localhost:8080

Application would take you to the login page. From there it self explanatory. Happy coding!

Navigate these URLs with various credentials listed on the login page. Observe the application logs to learn more.

  * http://localhost:8080/manage/users.html
  * http://localhost:8080/admin/dashboard.html
  * http://localhost:8080/login.html