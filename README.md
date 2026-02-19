## JetRoute (Microservice Gateway)
Purpose of this gateway is to route required request to the correct destination service.
To ensure that request redirect to the correct service, validation of the request by integrating authentication service, request logging to ensure our system is maintained and many more features.

### Features

- Reverse proxy -> Done
- Service Registry -> Done
- Authentication -> Done
- Request Logging -> Done
- Rate Limiting

### Workflow

```workflow
HTTP Request -> Gateway -> Middlewares -> Destination Service
```


### Configure Gateway
Create config.json file in root directory of application, information regarding the services should be described as mentioned below.

Please Note: If route is mentioned in the config then it is a private route which requires authentication, ensure authentication service is integrated if private route exists, by default all routes are public.


```config
{
  "service": "order-service",
  "host": "localhost",
  "port": 8090,
  "auth": {
    "host": "localhost",
    "port": 8090,
    "path": "/api/validate-token"
  },
  "private-routes": [
    {
      "path": "/api/orders/*"
    }
  ]
}
```