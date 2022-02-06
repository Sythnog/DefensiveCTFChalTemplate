# Use Alpine Linux, and set the working directory
FROM alpine:3.15.0
WORKDIR /opt/attVal

# Update the system, and install the needed packages
RUN apk update --no-cache && apk add bash go --no-cache

# Add a user for system-hardening
RUN adduser -D attValUser

# Move files to container, and build the application
COPY . .
RUN go build -race

# Cleanup, switch user, expose port, and start
RUN rm validateCode.sh main.go go.mod go.sum Dockerfile runAttacker.sh installValidatorTools.sh

# Switch user and install nuclei for this user
USER attValUser
RUN go install -v github.com/projectdiscovery/nuclei/v2/cmd/nuclei@latest
EXPOSE 80
ENTRYPOINT ["bash"]
CMD ["start.sh"]