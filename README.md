# Bee happy API

## Table of contents

- Provisions services

## Provisions services

- Provision mysql server with docker

```bash
docker run --name mysql \
-p 3306:3306 \
-e MYSQL_ROOT_PASSWORD=<mysql_your_custom_password> \
-d mysql:8.3.0
```
