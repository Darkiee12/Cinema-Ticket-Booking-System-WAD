FROM oven/bun:alpine
WORKDIR /home/bun/app
COPY package.json bun.lockb ./
RUN bun install
COPY . .
RUN bun run build
EXPOSE 5173
CMD ["bun", "run", "start"]
