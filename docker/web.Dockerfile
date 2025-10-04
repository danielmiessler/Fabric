FROM node:20-slim AS base
ENV PNPM_HOME=/pnpm
ENV PATH=$PNPM_HOME:$PATH
RUN corepack enable

FROM base AS deps
WORKDIR /app
COPY web/pnpm-lock.yaml web/package.json ./
RUN pnpm install --frozen-lockfile

FROM deps AS builder
WORKDIR /app
COPY web ./
RUN pnpm build

FROM base AS runtime
WORKDIR /app
ENV NODE_ENV=production
COPY --from=deps /app/node_modules ./node_modules
COPY web/package.json web/pnpm-lock.yaml ./
COPY --from=builder /app/build ./build
COPY web/svelte.config.js web/vite.config.ts web/tailwind.config.ts web/postcss.config.js web/app.html ./

EXPOSE 5173

CMD ["pnpm", "preview", "--host", "0.0.0.0", "--port", "5173"]
