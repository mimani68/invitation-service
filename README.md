# Invitation service 

## API

POST /console/login [login admin]
GET /console/codes
GET /console/codes/:codeId
POST /console/codes
POST /console/codes/:codeId/inactive
DELETE /console/codes/:codeId

POST /code/:code (login, retry new code)

## Development

Alter config file and replace sutable value for each settings.
```bash
vi config.yaml
```
Now you can start project
```bash
go run . serve http
```

## Production

```bash
go build -o ./bin/block_explorer .
```