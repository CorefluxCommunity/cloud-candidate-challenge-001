FROM golang:1.22-alpine3.20

WORKDIR /app

# install dependencies and terraform to execute terraform init
RUN apk add --no-cache wget unzip

RUN wget https://releases.hashicorp.com/terraform/1.4.6/terraform_1.4.6_linux_amd64.zip \
    && unzip terraform_1.4.6_linux_amd64.zip \
    && mv terraform /usr/local/bin/ \
    && rm terraform_1.4.6_linux_amd64.zip

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o app .

RUN chmod +x ./app

ENTRYPOINT ["./app"]

