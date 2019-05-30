FROM node:12.3-alpine

WORKDIR /var/app

COPY package.json /var/app

RUN yarn install

COPY . /var/app

EXPOSE 3000

CMD ["yarn", "run", "start"]
