# Invitation service 

## Specifications

| Title | Tools | Version |
|:--:|:--:|:--:|
| Language | Golang | 1.17.5 linux/amd64 |
| Framework | echo | 4 |
| Database | Mysql | 8 |
| Orm | sql native | * |
| Command line handler | urfave | * |
| Api documenation | postman | 2 |
| Container | Docker engine | 20.10.7 |
| Orchestartor | compose | 1.29.2 |

* REST-full APIs
* Separate concerns
* Fully clean code and inspired by clean architecture
* Admin endpoints protected with **basic auth**
* Full CRUD of *INVITATION*, inadition **DISABLE** feature for single code
* Invite tokens to expire after 7 days
* A public endpoint for validating the invite token
* Production ready by deployment instructions

## Development

Alter config file and replace suitable value for each settings.
```bash
vi config.yaml
```
Now you can start project
```bash
go run . serve http
```

## Production

If you want use binary file of server, fastest and easiest way which possible is generating machine readable binary file depend on your operating system.
```bash
go build -o ./bin/app .
```

or in cloud-native technology you can
```bash
docker-compose -p app -f deployment/docker-compose.yml up -d
```

## Documenation

See folder *api/* and file *pluseId.postman_collection.json*
