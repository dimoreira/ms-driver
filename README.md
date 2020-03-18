# Driver Microservice

Microservice that manages and returns info about Drivers

**Endpoints**:

- `GET /`
  - Returns the microservice basic info

- `GET /drivers`
  - Returns an list of drivers available

- `GET /drivers/{driverId}`
  - Returns a driver when searched by `uuid`

## How to run?

Just download the image and run in your machine!

```bash
$ docker run -it -p 8081:8081 dimoreira/ms-driver
```