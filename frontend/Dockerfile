# Use multi-stage builds to optimize the image size
FROM node AS builder

WORKDIR /app

# Copy only necessary files
COPY package.json package-lock.json ./
RUN npm install

COPY . . 
# Ensure assets are included in the build context


RUN npm run build

# Create a minimal runtime image
FROM node:alpine

WORKDIR /app

COPY --from=builder /app/.output .

ENTRYPOINT ["node", "server/index.mjs"]
