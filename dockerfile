FROM golang:latest

RUN apt-get update && apt-get install -y \
    wget \
    unzip \
    && rm -rf /var/lib/apt/lists/*

ENV TERRAFORM_VERSION=1.8.5

RUN wget https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
    unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
    mv terraform /usr/local/bin/ && \
    rm terraform_${TERRAFORM_VERSION}_linux_amd64.zip

RUN terraform --version

WORKDIR /app

COPY . .

WORKDIR /app/server

RUN go mod download
RUN go build -o /server

EXPOSE 8080
CMD [ "/server" ]
