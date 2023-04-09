# Go Microservices

- Front-end: `driver` to communicate with Broker-service
- Broker Service: `Mediator` between services
- Auth Service: Responsible for user authentication with `PostgresSQL`
- Logger Service: Responsible for logging info with `MongoDB`
- Mailer Service: Responsible for sending emails with `MailHog`

## How to start services

### Backend

```sh
make up_build 
```

### Frontend

```sh
make start 
```
