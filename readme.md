# Readme

## Migration

Install migration tool (MacOS)

```bash
brew install golang-migrate
```

Create migration file

```bash
migrate create -seq -ext=.sql -dir=./migrations create_users_table
```

## Run server

```bash
go run ./
```

## Ping

```bash
curl -i http://localhost:4000/ping

```

## Test basic authentication

Create some username password:

```bash
CREDENTIALS=$(echo -n "alice@example.com:pa55word" | base64)
```

Send the request

```bash
curl -i -H "Authorization: Basic $CREDENTIALS" http://localhost:4000/basic-auth
```
