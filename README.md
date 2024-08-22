# Go-buks API

A simple API to manage book lists, using Fiber and PostgreSQL.

[![Go](https://img.shields.io/badge/Demo_API-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go-buks.fly.dev/docs/)

## Tech Stacks

- **Web Framework**: [Fiber](https://docs.gofiber.io/).
- **Database**: [PostgreSQL](https://www.postgresql.org/).
- **ORM**: [GORM](https://gorm.io/)
- **Developments**: [Docker](https://www.docker.com/)

## Run Locally

Clone the project

```bash
git clone https://github.com/irfnd/go-buks
```

Go to the project directory

```bash
cd go-buks
```

Install dependencies

```bash
go mod download && go mod verify
```

Setup Database

```bash
cd database
```

```bash
docker-compose up -d
```

Copy .Env

```bash
cp .env.example .env
```

Start the server

```bash
go run .
```

## API Docs

You can view the API documentation at `http://localhost:8080/docs`

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are greatly appreciated.

- Fork the Project
- Create your Feature Branch `git checkout -b feature/AmazingFeature`
- Commit your Changes `git commit -m '[Add] - some AmazingFeature`
- Push to the Branch `git push origin feature/AmazingFeature`
- Open a Pull Request

## License

Distributed under the [MIT](https://github.com/irfnd/go-buks/blob/master/LICENSE) License.
