FROM oven/bun:alpine
WORKDIR /home/bun/app
COPY package.json bun.lockb ./
RUN bun install
COPY . .
EXPOSE 80

RUN bun install
CMD ["bun", "run", "dev"]