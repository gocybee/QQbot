# To train rasa on my stupid mac M1 chip
FROM --platform=amd64 python:3.8-slim-buster

RUN set -x ; pip3 install -U --user pip && pip3 install rasa

# Valid installation
RUN rasa --version

CMD 'bash'