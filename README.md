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


<details>
<summary>Writing new API</summary>

> Goal: Add a new `GET /v1/products` endpoint


1. Go to `rest/api/v1` folder
2. Create a new file `product_controller.go`
3. Create a new struct `ProductController`
4. Create a constructor `NewProductController`
5. Add a method `List`

```go
type ProductController struct {
	productUC ProductUsecase
}

func (h *ProductController) List(w http.ResponseWriter, r *http.Request) {
	p, err := h.productUC.List(r.Context())
	if err != nil {
		response.JSONError(w, err)
		return
	}

	response.JSON(w, response.OK(&p), http.StatusOK)
}
```

6. Go to `rest/api/v1.go`
7. Add the `ProductController` to the `API` struct
8. Mount the routes accordingly

```go
type API struct {
	*ProductController
}

func (api *API) Register(r chi.Router) {
	r.Route("/v1", func(r chi.Router) {
		r.Route("/products", func(r chi.Router) {
			r.Get("/", api.ProductController.List)
		})
	})
}
```

</details>

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
