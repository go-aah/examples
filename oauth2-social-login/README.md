<p align="center">
  <img src="https://cdn.aahframework.org/assets/img/aah-logo-64x64.png" />
  <h2 align="center">Example - OAuth2 Social Login</h2>
</p>

This example demonstrates OAuth2 3-legged Flow implementation with aah framework. It does social login with Facebook, Google and Github.

Learn more about [OAuth2 Auth scheme](https://docs.aahframework.org/auth-schemes/oauth2.html)

## How to get the aah examples?

```bash
git clone https://github.com/go-aah/examples.git $GOPATH/src/aahframework.org/examples
```

## How to run this example?

**Step 1:**

First, you have to configure the OAuth2 provider(s) Client ID and Secret in the `security.conf`. If you want to try only one provider, just configure them alone.

**Step 2:**

```bash
aah r -i aahframework.org/examples/oauth2-social-login
```

### Now visit this URL

  * http://localhost:8080

On application home page click on the provider icon :smile:.