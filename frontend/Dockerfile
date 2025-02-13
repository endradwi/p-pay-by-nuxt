# Use multi-stage builds to optimize the image size
FROM node AS builder

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install

COPY . . 

RUN npm run build

FROM node:alpine

WORKDIR /app

COPY --from=builder /app/.output .

ENTRYPOINT ["node", "server/index.mjs"]
