# go-api-test



## How to structure your APIs

- structure the folders based on the API routes
	- e.g. /health -> /rest/api/health_controller.go
	- e.g. /v1/products -> /rest/api/v1/product_controller.go
- separate the handler from the framework, this allows you to change routers
- add graceful shutdown
- setup read and write timeouts
- limit the buffer size
- copy the body, so that you can retrieve it later for logging purposes
- map domain errors to http status errors
- document the environment variables in the `.env.sample`

```.env
APP_VERSION=<optional: the current app version, e.g. 0.0.1>
JWT_SECRET=<required: provide a secret for jwt>
```
