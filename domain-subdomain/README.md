<p align="center">
  <img src="https://cdn.aahframework.org/assets/img/aah-logo-64x64.png" />
  <h2 align="center">Example - Domain, Subdomain and Wildcard Subdomain</h2>
</p>

This example application gives insights on how to use domains, subdomains and wildcard subdomains with aah framework. Reference to [Routes Config](https://docs.aahframework.org/routing.html).

### Get aah examples

```bash
git clone https://github.com/go-aah/examples.git $GOPATH/src/aahframework.org/examples
```

### Configure local DNS mapping

Firstly, configure `hosts` file. The purpose is to do local DNS mapping for `sample.com`. This step is not applicable for production. Typically, these settings happen in domain DNS manager. Learn more about [Wildcard DNS - wikipedia](https://en.wikipedia.org/wiki/Wildcard_DNS_record)

This [rackspace article](https://support.rackspace.com/how-to/modify-your-hosts-file/) covers the steps to modify-your-hosts-file for Mac, Linux and Windows.

```bash
127.0.0.1       sample.com admin.sample.com username1.sample.com username2.sample.com username3.sample.com
```

### Run this example

```bash
aah r -i aahframework.org/examples/domain-subdomain
```

### Visit these URLs

  * http://sample.com:8080
  * http://admin.sample.com:8080
  * http://username1.sample.com:8080
  * http://username2.sample.com:8080
  * http://username3.sample.com:8080

