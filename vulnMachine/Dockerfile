FROM debian:stable-slim
WORKDIR /opt/notFoundService
COPY . .
RUN bash install.sh
RUN groupadd -r server && useradd -r -g server server
USER server
ENTRYPOINT ["python3"]
CMD ["vulnService.py"]