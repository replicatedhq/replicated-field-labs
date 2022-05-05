FROM ubuntu:18.04

MAINTAINER Fernando Cremer "cremerfc@gmail.com"

RUN apt-get update -y && \
    apt-get install -y python3-pip python3-dev

COPY ./Requirements.txt /Requirements.txt

WORKDIR /

RUN pip3 install -r Requirements.txt

COPY . /

ENTRYPOINT [ "python3" ]

CMD [ "app/app.py" ]