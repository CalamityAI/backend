
# Backend
Documentation

## Quick start
```
docker-compose up
```
The app should be available here: http://localhost:5000

Api docs can be found here: http://localhost:5000/api-docs/index.html

## Development

### Preparing for work

Install docker

Install dependencies
```bash
  go mod download
```

Install atlas (for database migrations) (this works only for linux/mac. If you are using windows, please see the official docs)
```bash
curl -sSf https://atlasgo.sh | sh
```

Install air (for hot reloading)
```bash
go install github.com/air-verse/air@latest
```

Configure golangci-lint (linter) in your IDE. See https://golangci-lint.run/welcome/integrations/

Duplicate `.example.dev` file and rename the copy to `.env`

To run the development database run this
```bash
docker-compose up database
```
or run database locally and fill `.env` with db connection data (not recommended)

### Running
Run with hot reloading
```bash
air
```

Run
```
go run cmd/http/main.go
```

### Database management

Create migration

```bash
```

Apply migration

```bash
```

### Linting

Lint
```bash
```

### Testing

Run tests
```bash
```


