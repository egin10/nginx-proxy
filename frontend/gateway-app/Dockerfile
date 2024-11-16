# Step 1: Build the React app
FROM node:18 as build

WORKDIR /app

# Copy the package.json and install dependencies
COPY package.json ./
RUN npm install

# Copy the rest of the app's source code and build it
COPY . .
RUN npm run build

# Step 2: Serve the React app with nginx
FROM nginx:alpine

# Copy the built app from the previous stage
COPY --from=build /app/dist /usr/share/nginx/html
COPY --from=build /app/.htaccess /usr/share/nginx/html/dist

# Copy the custom nginx configuration file
COPY --from=build /app/nginx.conf /etc/nginx/conf.d/default.conf

# Expose port 80
EXPOSE 80

# Start nginx server
CMD ["nginx", "-g", "daemon off;"]
