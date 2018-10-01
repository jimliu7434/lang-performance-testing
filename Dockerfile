FROM node:8

# Create app folder
RUN mkdir -p /var/app
WORKDIR /var/app

# Copy app files into app folder
COPY . /var/app

RUN ["yarn", "install", "--production"]

CMD ["yarn", "start"]
