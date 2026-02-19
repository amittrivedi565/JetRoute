## JetRoute (Microservice Gateway)
Purpose of this gateway is to route required request to the correct destination service.
To ensure that request redirect to the correct service, validation of the request by integrating authentication service, request logging to ensure our system is maintained and many more features.

### Features

- Reverse proxy
- Service Registry
- Request Validation
- Rate Limiting
- Request Logging
- Circuit Breaker

### Workflow

```workflow
HTTP Request -> Gateway -> Middlewares -> Destination Service
```

- Register the service with gateway
- Integrate authentication service