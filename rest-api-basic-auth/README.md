<p align="center">
  <img src="https://cdn.aahframework.org/assets/img/aah-logo-64x64.png" />
  <h2 align="center">Example - REST API Basic Auth</h2>
</p>

This example demonstrates REST API Basic Auth with aah framework. That includes Authentication and Route Authorization via routes config. 

aah supports basic auth mechanism in two ways, choose per use case -

Realm | Description
----- | -----------
File | When you know set of pre-defiend subjects (aka users), roles and permissions (roles and permissions values are optional though)
Dynamic | Subject information lies in Data Source (DB, API provider, etc). Implementing interfaces `authc.Authenticator` and `authz.Authorizer`

Learn more about [Security design](https://docs.aahframework.org/security-design.html), [Authentication](https://docs.aahframework.org/authentication.html) and [Authorization](https://docs.aahframework.org/authorization.html).

## How to get the aah examples?

```bash
git clone https://github.com/go-aah/examples.git $GOPATH/src/aahframework.org/examples
```

## How to run this example?

```bash
aah r -i aahframework.org/examples/rest-api-basic-auth
```

## Use your favorite REST client to make a request

**Demo User Credentials**

Username/Password | Roles & Permissions | IsLocked
----------------- | ------------------- | --------
user1@aahframework.org/welcome123 | Roles: "employee", "manager" <br> Permissions: "user:read,edit:reportee" | No
user2@aahframework.org/welcome123 | Roles: "employee" <br> Permissions: N/A | No
user3@aahframework.org/welcome123 | Roles: "employee" <br> Permissions: N/A | Yes
admin@aahframework.org/welcome123 | Roles: "employee", "manager", "admin" <br> Permissions: "user:read,edit,delete:reportee" | No

<br>

**API Endpoints**

```
GET     /                       - Shows welcome message (Anonymous access)
GET     /v1/reportee/:email     - Returns user data for given email address based on authorization (Secured)
```

### Welcome Message

Send GET request to http://localhost:8080/

Response:

```json
{"message":"Welcome to aah framework - REST API Basic Auth Example"}
```

### Get reportee data

Send GET request to http://localhost:8080/v1/reportee/user1@aahframework.org with Basic Auth `user1@aahframework.org/welcome123`

Resposne:

```json
{
    "first_name": "East",
    "last_name": "Corner",
    "email": "user1@aahframework.org",
    "is_locked": false,
    "is_expried": false,
    "roles": [
        "employee",
        "manager"
    ],
    "permission": [
        "user:read,edit:reportee"
    ]
}
```

Now, try yourself with various combinations with above demo credentials.