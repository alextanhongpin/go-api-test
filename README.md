# go-api-test


## Server configuration

- separate the handler from the framework, this allows you to change routers
- add graceful shutdown
- setup read and write timeouts
- limit the buffer size

## How to structure your APIs

- structure the folders based on the API routes
	- e.g. /health -> /rest/api/health_controller.go
	- e.g. /v1/products -> /rest/api/v1/product_controller.go
- middlewares, context, and error
- serialization for response (data envelope)
- copy the body, so that you can retrieve it later for logging purposes
- map domain errors to http status errors
- document the environment variables in the `.env.sample`


## Authorization

- how to handle auth
- getting jwt claims

## Request/response

- validation request
- request payload size
- body parser
- trim strings
- query filters
- url builders
- forms and file uploads
- response envelope, links, status code and error handling

## Middlewares

- request id
- cors
- auth bearer/basic
- healthcheck

```.env
APP_VERSION=<optional: the current app version, e.g. 0.0.1>
JWT_SECRET=<required: provide a secret for jwt>
```


## How to contribute

- first endpoint
- document endpoint
- first test
- conventions


## Advanced

- authorization
- whitelist ip
- webhooks (notifications, callbacks) handler, security and testing
- localization
- versioning
- dependency injection
- OTP flow

## Generating Token


Call the register endpoint to generate the token.
```bash
# -r means raw output. We want the string without the json double quotes
# Sends the output to the clipboard.
$ curl -XPOST localhost:8080/register | jq -r .data.accessToken | pbcopy
```

Make a call to the protected endpoint using the token:

```bash
$ curl -XPOST -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODI1MjgxODMsInN1YiI6IjllZTNkZDI2LWY5MWItNDNjMy04NzJkLTJlNjg0YzBjOTIzYyJ9.GFZl5v0JXC72PpGa2953Ioh3xd7nM9ezI4YL-rYNK7Q' localhost:8080/v1/categories
```
