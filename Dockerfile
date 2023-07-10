FROM --platform=x86_64 golang

WORKDIR /app

COPY . .

EXPOSE 9000

CMD ["go", "run", "main.go"]
