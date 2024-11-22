# Go API Server for location-api

The ETSI MEC ISG MEC013 Location API described using OpenAPI.

## Overview
This server was generated by the [openapi-generator]
(https://openapi-generator.tech) project.
By using the [OpenAPI-Spec](https://github.com/OAI/OpenAPI-Specification) from a remote server, you can easily generate a server stub.
-

To see how to make this your own, look here:

[README](https://openapi-generator.tech)

- API version: 3.1.1
- Build date: 2024-09-19T17:42:26.928041535+02:00[Europe/Berlin]
- Generator version: 7.8.0
For more information, please visit [https://forge.etsi.org/rep/mec/gs013-location-api](https://forge.etsi.org/rep/mec/gs013-location-api)

### Running the server

To run the server, follow these simple steps:

```
go run main.go
```

To run the server in a docker container
```
docker build --network=host -t location-api .
```

Once the image is built, just run
```
docker run --rm -it location-api
```

### Known Issue

Endpoints sharing a common path may result in issues. For example, `/v2/pet/findByTags` and `/v2/pet/:petId` will result in an issue with the Gin framework. For more information about this known limitation, please refer to [gin-gonic/gin#388](https://github.com/gin-gonic/gin/issues/388) for more information.

A workaround is to manually update the path and handler. Please refer to [gin-gonic/gin/issues/205#issuecomment-296155497](https://github.com/gin-gonic/gin/issues/205#issuecomment-296155497) for more information.