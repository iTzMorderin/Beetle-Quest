# Beetle Quest

A beautifull and fullfilling web based gacha system.

## Get Started

This guide will help you set up and run the application using Docker and Docker Compose. By following these steps, you'll have a local version of the application running on your machine.
Prerequisites:

Before you begin, ensure you have the following installed:

-   Docker
-   Docker Compose

### Cloning the repo

Clone the repository to your local machine using Git.

```bash
git clone https://github.com/iTzMorderin/Beetle-Quest
cd Beetle-Quest
```

### Configuring the application

The application uses environment variables to configure itself. Their values can be modified inside the docker compose file located in the deploy directory. Instead the `secrets` values are stored inside inside `secrets/` folder with the following structure:

```bash
deploy/
├──secrets/
│  ├── jwt.env
│  ├── oauth2.env
│  ├── postgress.env
│  └── redis.env
...
└──compose.yml
```

### Starting the application

To start the application, run the following commands:

```bash
cd beetle-quest/deploy
docker compose up
```

If you want to run the application in the background, you can use the `-d` flag:

```bash
docker compose up -d
```

> [!NOTE]
> The service gui is available at `https://localhost/static`.

> [!WARNING]
> The admin interface is not available, but it can be accessed with tools like `curl`. The admin's
> login requires `otp`, the seed can be retrieved from the qrcode inside `assets/admin_otp_qrcode.png`.
> Admins CANNOT be created, the credentials for the development one are:
>
> -   AdminID: `09087f45-5209-4efa-85bd-761562a6df53`
> -   Password: `admin`
> -   OTP: retrieved from the qrcode

An example request to login as an admin:

```bash
curl -v -k -X POST \
  -H "Content-Type: application/json" \
  -d '{"admin_id": "09087f45-5209-4efa-85bd-761562a6df53", "password": "admin", "otp_code": "<OTP_CODE>"}' \
  https://localhost/api/v1/auth/admin/login
```

### Stopping the application

To stop the application, run the following command:

```bash
cd beetle-quest/deploy
docker compose down
```

## Tests

> [!IMPORTANT]
> The tests needs to be executed in a clean environment. If the system has been used before we recommend using a
> a `docker compose down -v` to remove all volumes and start from scratch.

### Postman

You fill find the Postman collection file`collection.json` inside `beetle-quest/tests/postman/`, you can execute them with Postam Newman:

```sh
docker run --rm --net beetle-quest_internal -v <path/to/this/repo>/beetle-quest/tests/postman/collection.json:/collection.json postman/newman run /collection.json --insecure --ignore-redirects --color on
```

### Locust

Locust's tests can be found inside `beetle-quest/tests/locust/` folder, to start locust, and run the tests, these commands have to be executed:

```sh
cd beetle-quest/tests/locust/
docker build -t beetle-quest-locust:latest .
docker run --rm --network=beetle-quest_external -p 127.0.0.1:8089:8089 beetle-quest-locust:latest
```

## Project structure

The project is structured as follows:

```bash
.
├── assets/
│   ├── drawio
│   ├── gimp
│   └── images
├── beetle-quest/
│   ├── api/
│   ├── cmd/
│   ├── deploy/
│   ├── go.mod
│   ├── go.sum
│   ├── internal/
│   ├── pkg/
│   └── templates/
├── doc/
└── README.md
```

## References

-   [Project Structure 1](https://betterprogramming.pub/how-are-you-structuring-your-go-microservices-a355d6293932)
-   [Project Structure 2](https://gochronicles.com/project-structure/)
