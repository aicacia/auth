# Auth API

### Deploy

- `docker build -t ghcr.io/aicacia/auth-api:latest .`
- `docker push ghcr.io/aicacia/auth-api:latest`
- `helm upgrade auth-api helm/auth-api -n api --install -f values.yaml --set image.hash="$(docker inspect --format='{{index .Id}}' ghcr.io/aicacia/auth-api:latest)"`

### Undeploy

- `helm delete -n api auth-api`
