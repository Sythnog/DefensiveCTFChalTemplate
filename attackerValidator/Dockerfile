# Use latest version of stable debian, and set working folder in the container
FROM debian:stable-slim
WORKDIR /opt/attackerValidator

# Install needed packages
RUN apt update && apt install python3 python3-pip -y

# Move files and install python dependencies
COPY . .
RUN pip3 install -r requirements.txt

# Add a user to not run the service as root, and switch to it
RUN groupadd -r server && useradd -r -g server server
USER server

# Expose the used port
EXPOSE 80

# Start the service
ENTRYPOINT ["bash"]
CMD ["start.sh"]