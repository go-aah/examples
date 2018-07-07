<p align="center">
  <img src="https://cdn.aahframework.org/assets/img/aah-logo-64x64.png" />
  <h2 align="center">Example - REST API Services</h2>
</p>

This example demonstrates buliding REST API services with aah framework. It gives an overview of aah features and API capabilities.

Learn more about [REST API versioning strategies](https://docs.aahframework.org/rest-api-versioning.html) and [Auto parse and Bind](https://docs.aahframework.org/request-parameters-auto-bind.html).

### Get aah examples

```bash
git clone https://github.com/go-aah/examples.git $GOPATH/src/aahframework.org/examples
```

### Run this example

```bash
aah r -i aahframework.org/examples/rest-api
```

### Use a preferred REST client to make a request

Example application has endpoints with in-memory blog post implementation. Observe application logs to learn more.

```
GET     /              - Welcome message
POST    /v1/posts      - Creates a simple blog post with `Title` and `Body`
GET     /v1/posts/:id  - Get blog post content by ID
PUT     /v1/posts/:id  - Update blog post content by ID with `Title` and `Body`
DELETE  /v1/posts/:id  - Delete blog post by ID
GET     /v1/posts      - Get all the blog posts
```

### Welcome Message

Send GET request to http://localhost:8080/

Response:

```json
{"message": "Welcome to aah framework - REST API Services example"}
```

### Create a Blog Post

Send POST request to http://localhost:8080/v1/posts

Header: Content-Type -> application/json or text/json

Request Body:

```json
{
    "title": "Hey, my first post",
    "body": "<p>This is my first blog post</p>"
}
```

### Get Blog Post by ID

Send GET request to http://localhost:8080/v1/posts/1

Response:

```json
{
    "id": 1,
    "title": "Hey, my first post",
    "body": "<p>This is my first blog post</p>",
    "created_at": "2017-08-31T13:10:50-07:00",
    "updated_at": "2017-08-31T13:10:50-07:00"
}
```

### Get All Blog Posts

Send GET request to http://localhost:8080/v1/posts

Response:

```json
{
    "posts": [
        {
            "id": 1,
            "title": "Hey, my first post",
            "body": "<p>This is my first blog post</p>",
            "created_at": "2017-08-31T13:10:50-07:00",
            "updated_at": "2017-08-31T13:14:15-07:00"
        }
    ]
}
```

### Update Blog Post by ID

Send PUT request to http://localhost:8080/v1/posts/1

Header: Content-Type -> application/json or text/json

Request Body:

```json
{
    "title": "Hey, my first post - update",
    "body": "<p>This is my first blog post - update</p>"
}
```

Response: 204 No Content

### Delete Blog Post by ID

Send PUT request to http://localhost:8080/v1/posts/1

Response: 204 No Content
