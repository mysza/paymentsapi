# Payments API

Payments API is an examplary HTTP RESTful service.

## Code organisation

-   /api - HTTP handlers for all the methods exposed
-   /cmd - the entry point to the appliaction (setting up the service)
-   /domain - the domain model
-   /repository - the repository of the service
-   /service - core implementation of the service functionality

## Building

Run `make build`

## Testing

Run `make test` to run the test, `run testcover` to create and open test coverage report.

## Running

Run `make run` - the build will be triggered before the start of the service.

`PAYMENTSAPI_PORT` environment variable can be used to define on which port shal the HTTP server listen (default is
3000).

`PAYMENTSAPI_DBDIR` environment variable can be used to define the directory where the database files will be stored.
Default is `./db`.

## License

MIT
