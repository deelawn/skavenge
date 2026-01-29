# Simple Dockerfile for Skavenge Landing Page
# Serves static HTML/CSS/JS on port 80 using nginx

FROM nginx:alpine

# Copy the landing page to nginx html directory
COPY index.html /usr/share/nginx/html/

# Expose port 80
EXPOSE 80

# Nginx starts automatically
