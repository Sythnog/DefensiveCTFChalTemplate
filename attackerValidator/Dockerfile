FROM debian:stable-slim
WORKDIR /opt/attackerValidator
COPY . .
RUN bash install.sh
RUN groupadd -r server && useradd -r -g server server
USER server
ENTRYPOINT ["bash"]
CMD ["start.sh"]