FROM golang:1.23

# Establece el directorio de trabajo
WORKDIR /go/src/app

# Copia el archivo go.mod y go.sum
COPY go.mod go.sum ./

# Descarga las dependencias
RUN go mod tidy

# Copia el código fuente
COPY . .

# Compila la aplicación Go
RUN go build -o app .

# Expone el puerto en el contenedor
EXPOSE 8090

# Comando para ejecutar la aplicación
CMD ["./app"]