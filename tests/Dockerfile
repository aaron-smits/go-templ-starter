# Get the latest version of Playwright
FROM mcr.microsoft.com/playwright:v1.41.1-jammy
 
# Set the work directory for the application
WORKDIR /app
 
# Set the environment path to node_modules/.bin
ENV PATH /app/node_modules/.bin:$PATH

ENV BASE_URL=http://localhost:8080

# COPY the needed files to the app folder in Docker image
COPY . .

RUN npm i

