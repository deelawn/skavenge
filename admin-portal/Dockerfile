# Build stage
FROM node:20-alpine AS build

WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm ci

# Copy source code
COPY . .

# Build the application
RUN npm run build

# Production stage
FROM node:20-alpine

WORKDIR /app

# Copy package files
COPY package*.json ./

# Install production dependencies only
RUN npm ci --production

# Copy built assets from build stage
COPY --from=build /app/dist ./dist
COPY --from=build /app/public ./public

# Expose port
EXPOSE 3000

# Start the application
CMD ["npm", "run", "serve"]
