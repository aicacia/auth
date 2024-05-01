# Auth UI

## Tools

- [nvm](https://github.com/nvm-sh/nvm#installing-and-updating)
- [pnpm](https://pnpm.io/installation)
- [tailwind css](https://tailwindcss.com/docs)
- [svelte kit](https://kit.svelte.dev/docs)
- [icons](https://lucide.dev/icons/)

## Docker/Helm

### Deploy

- `docker build --build-arg PUBLIC_AUTH_API_URL=https://api.auth.aicacia.com -t ghcr.io/aicacia/auth-ui:latest .`
- `docker push ghcr.io/aicacia/auth-ui:latest`
- `helm upgrade auth-ui helm/auth-ui -n ui --install -f values.yaml --set image.hash="$(docker inspect --format='{{index .Id}}' ghcr.io/aicacia/auth-ui:latest)"`

### Undeploy

- `helm delete -n ui auth-ui`
