# Multi-stage Dockerfile for Skavenge React Web App
# Stage 1: Build the React application

FROM node:18-alpine AS build

WORKDIR /app

# Copy package files
COPY package.json package-lock.json* ./

# Install dependencies
RUN npm install

# Copy source code
COPY public/ ./public/
COPY src/ ./src/

# Build the React app
RUN npm run build

# Stage 2: Serve with NGINX

FROM nginx:alpine

# Copy built React app from build stage
COPY --from=build /app/build /usr/share/nginx/html

# Create a default config.json file (will be overridden by volume mount)
RUN echo '{"contractAddress":"0x0000000000000000000000000000000000000000","networkRpcUrl":"http://localhost:8545"}' > /usr/share/nginx/html/config.json

# Expose port 80
EXPOSE 80

# Nginx will start automatically
