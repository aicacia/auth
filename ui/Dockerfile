FROM --platform=linux/amd64 node:22-alpine3.19 as node-builder
RUN npm install -g pnpm@9

WORKDIR /app

COPY package.json pnpm-lock.yaml ./
RUN pnpm install

ARG PUBLIC_URL=https://auth.aicacia.com
ENV PUBLIC_URL=${PUBLIC_URL}
ARG PUBLIC_AUTH_API_URL=https://api.auth.aicacia.com
ENV PUBLIC_AUTH_API_URL=${PUBLIC_AUTH_API_URL}
ARG NODE_ENV=production
ENV NODE_ENV=${NODE_ENV}

RUN echo "PUBLIC_URL=${PUBLIC_URL}" >> .env
RUN echo "PUBLIC_AUTH_API_URL=${PUBLIC_AUTH_API_URL}" >> .env

COPY src ./src
COPY static ./static
COPY postcss.config.cjs svelte.config.js tailwind.config.cjs vite.config.js ./

RUN pnpm run build

FROM --platform=linux/amd64 nginx:1.27-alpine3.19-slim
LABEL org.opencontainers.image.source https://github.com/aicacia/auth

ARG MAINTENANCE_MODE=false
ENV MAINTENANCE_MODE=${MAINTENANCE_MODE}

COPY default.conf.template /etc/nginx/templates/default.conf.template
RUN rm -rf /usr/share/nginx/html
COPY 40-env.sh /docker-entrypoint.d/40-env.sh

COPY --from=node-builder /app/build /usr/share/nginx/html
