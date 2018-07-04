<p align="center">
  <img src="https://cdn.aahframework.org/assets/img/aah-logo-64x64.png" />
  <h2 align="center">Example - REST API JWT Auth</h2>
</p>

This example demonstrates REST API JWT Auth with aah framework. aah REST API JWT Auth includes authentication and route authorization via routes config.

Generic Auth can be customized in so many ways. This application implements JSON Web Token (JWT) using Generic auth scheme.

Learn more about [Security design](https://docs.aahframework.org/security-design.html), [Authentication](https://docs.aahframework.org/authentication.html) and [Authorization](https://docs.aahframework.org/authorization.html).

### Get aah examples

```bash
git clone https://github.com/go-aah/examples.git $GOPATH/src/aahframework.org/examples
```

### Run this example

```bash
aah r -i aahframework.org/examples/rest-api-jwt-auth
```

### Use a preferred REST client to make a request

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
POST    /v1/token               - Issues JSON Web Token (JWT) for given username and password (Anonymous access)
GET     /v1/reportee/:email     - Returns user data for given email address based on authorization (Secured)
```

### Welcome Message

Send GET request to http://localhost:8080/

Response:

```json
{"message":"Welcome to aah framework - REST API Basic Auth Example"}
```

### Get Access Token (JWT)

Send POST request to http://localhost:8080//v1/token

Request Body:

```json
{
  "username": "user1@aahframework.org",
  "password": "welcome123"
}
```

Response:

Returns access token for valid credentials.

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzA1NzY3MzcsInVzZXJuYW1lIjoidXNlcjFAYWFoZnJhbWV3b3JrLm9yZyJ9.944bfZpGY8I4ktJzKPA6pJFjhIW53upQBlVT7xSJwPA",
    "token_type": "bearer"
}
```

### Get Reportee data

Send GET request to http://localhost:8080/v1/reportee/user2@aahframework.org

Pass access token via HTTP header `Authorization: Bearer <access-token>`

Resposne:

```json
{
    "first_name": "West",
    "last_name": "Corner",
    "email": "user2@aahframework.org",
    "is_locked": false,
    "is_expried": false,
    "roles": [
        "employee"
    ]
}
```

Now, try various combinations with the above demo credentials.
