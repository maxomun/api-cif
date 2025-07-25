# Usa una imagen base oficial de Go para la construcción
FROM golang:1.24.4 as builder

RUN apt-get update && apt-get install -y ca-certificates openssl

ARG cert_location=/usr/local/share/ca-certificates

# Get certificate from "github.com"
RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.crt
# Get certificate from "proxy.golang.org"
RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${cert_location}/proxy.golang.crt
# Update certificates
RUN update-ca-certificates

ENV GIT_SSL_NO_VERIFY=true

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos go.mod y go.sum, para instalar las dependencias primero (esto ayuda con el caching)
COPY go.mod go.sum ./

# Descarga todas las dependencias del proyecto
RUN go mod download

# Copia el código fuente del proyecto
COPY . .

# Compila la aplicación
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Usa una imagen base mínima para ejecutar el binario compilado
FROM alpine:latest  

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /root/

# Copia el binario compilado desde la etapa de construcción anterior
COPY --from=builder /app/main .

# Expone el puerto 8080
EXPOSE 8080

# Comando que se ejecuta cuando se inicia el contenedor
CMD ["./main"]
