# Liatrio Demo 

---
## What we will be going over:

- Introduction
- Go over main.go go.mod and go.sum
- Local testing verify it works locally by making curl requests on localhost
- Go over complications of testing and commit versions are important
- Secrets for github actions, and why it is important?
- Explain what the Dockerfile is doing
- Explain build.yml

## Commands

```bash

# 2. Build fresh image
docker build -t liatrio-demo .

# 3. Run locally on port 80
docker run --rm -d -p 80:80 --name liatrio-demo-test liatrio-demo

# 4. test & inspect logs
curl http://127.0.0.1:80/
docker logs liatrio-demo-test

```

---

