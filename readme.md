# Authentication

## Run server

```bash
go run ./
```

## Ping

Test if server is working

```bash
curl -i http://localhost:4000/ping
---
pong
```

## Test basic authentication

Create some username password:

```bash
CREDENTIALS=$(echo -n "alice@example.com:pa55word" | base64)
```

You can double check if it worked:

```bash
echo $CREDENTIALS | base64 --decode
---
alice@example.com:pa55word                                                                     
```

Send the request

```bash
curl -i -H "Authorization: Basic $CREDENTIALS" http://localhost:4000/basic-auth
```
