FROM ubuntu:latest
LABEL authors="luismiranda"

WORKDIR /app

RUN apt-get update && \
    apt-get install -y wget unzip

RUN wget https://releases.hashicorp.com/terraform/1.0.0/terraform_1.0.0_linux_amd64.zip && \
    unzip terraform_1.0.0_linux_amd64.zip && \
    mv terraform /usr/local/bin/ && \
    rm terraform_1.0.0_linux_amd64.zip

COPY . /app

RUN chmod +x run.sh

CMD ["./run.sh"]
