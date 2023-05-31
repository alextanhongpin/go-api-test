# go-api-test


## Server Configuration

There are several best practices when it comes to configuring your server:

- separate the handler from the framework, this allows you to change routers
- add graceful shutdown
- set read timeouts
- set write timeouts
- handle context cancellation
- limit the size of the request payload and headers
- add rate-limiting for routes

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
<summary>Base Endpoint</summary>

### The API Struct
Each versioned endpoint will have an `API` struct. The root `/` endpoint `API` struct can be found in `rest/api/api.go`:

https://github.com/alextanhongpin/go-api-test/blob/9e8f4d96d543a98f71652da495177ad8664b8ff5/rest/api/api.go#L8-L12

Here, we register the resource controllers as well as middlewares for the endpoint. 
</details>

<details>
<summary>Adding Routes</summary>

## Adding Routes

Each `API` struct will have a `Register` method where we will register the resource routes.

https://github.com/alextanhongpin/go-api-test/blob/9e8f4d96d543a98f71652da495177ad8664b8ff5/rest/api/api.go#L14-L24
</details>

<details>
<summary>What are Controllers</summary>

## Controllers

Controllers are a collection of resources. Each controller can have several methods that maps to the HTTP methods.

</details>


<details>
<summary>Adding new API</summary>

This example demonstrates on how to add a new API endpoint

> Goal: Add a new `GET /v1/products` endpoint

1. Go to `rest/api/v1` folder
2. Create a new file `product_controller.go`
3. Create a new struct `ProductController`
4. Create a constructor `NewProductController`
5. Add a method `List`

```go
package v1

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
package v1

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

<details>
<summary>Guarding Routes</summary>

---
To guard routes, we can mount the `RequireAuth` middleware.

1. Go to `rest/api/v1.go` (or specific versioned endpoint)
2. Add the `RequireAuth` middleware to the struct `API`
3. Attach the `RequireAuth` to the routes that you want to protect in the `Register` method

```go
package v1

import (
	"github.com/alextanhongpin/core/http/middleware"
	"github.com/go-chi/chi/v5"
)

type API struct {
	RequireAuth middleware.Middleware
	*CategoryController
}

func (api *API) Register(r chi.Router) {
	r.Route("/v1", func(r chi.Router) {
		r.Route("/categories", func(r chi.Router) {
			// Attach to a single route
			r.With(api.RequireAuth).Post("/", api.CategoryController.Create)
		})
		
		// Attach to a group
		r.Group(func(r chi.Router) {
			r.Use(api.RequireAuth)
		})
	})
}
```


</details>

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

## Testing 

What is the goal of testing the API? There are many different ways of testing too, such as using Postman, writing code etc. 

For now, we stick with the following goals:

- tests as documentation guide
- tests as validation for behavior
- tests as a way to describe expected output json
- tests as a workflow guide

Test should serve as documentation. Tools like openAPI for example may show the sample expected request/response, but they don't show scenarios when you use different payload. For example, if you were to build a simple payment endpoint similar to Stripe, you will have different test cards that could trigger different scenarios. The requests are usually query string, path params, and body payload and http headers as well. The response we want to validate is usually the http headers as well as the payload body or error.

Test scenarios can be written in BDD style:

```markdown
Given that User calls the POST /payments
When the card is invalid
Then the API will error with status 422
And User will see Error Card Rejected.
```

Some business flows are easier to capture programmatically too. APIs workflows can consists of different steps, such as initially authenticating the users, then populating the data to be queries etc, as well as chaining multiple api steps.

Should the API be making actual database calls or mutating data? probably not. we just want to simulate the request response. 


