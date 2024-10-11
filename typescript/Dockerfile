FROM node:21

WORKDIR /app

COPY package.json tsconfig.json ./
RUN npm install

COPY flatten_directory.ts ./

RUN npm run build

ENTRYPOINT ["node", "dist/flatten_directory.js"]
