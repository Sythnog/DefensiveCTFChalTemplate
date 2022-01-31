# Use latest version of stable debian, and set working folder in the container
FROM debian:stable-slim
WORKDIR /opt/notFoundService

# Ensure updated apt, and install basic debug tools for users to utilize
RUN apt update && apt install tcpdump vim nano emacs htop wget curl python3 python3-pip python3-dev git libssl-dev libffi-dev build-essential -y

# Install useful python packages for debugging 
RUN pip3 install requests

# Setup SSH
RUN apt install -y openssh-server && mkdir /var/run/sshd && echo "root:notS0S3cr3tSSHPa55" | chpasswd && sed -ri "s/^#?PermitRootLogin\s+.*/PermitRootLogin yes/" /etc/ssh/sshd_config && sed -ri "s/UsePAM yes/#UsePAM yes/g" /etc/ssh/sshd_config && mkdir /root/.ssh

# Move files from local machine to container, run install script, and cleanup afterwards
COPY . .
RUN bash install.sh && rm install.sh runVulnMachine.sh Dockerfile

# Expose the used ports
EXPOSE 22
EXPOSE 80

# Start the machine by running "bash start.sh"
ENTRYPOINT ["bash"]
CMD ["start.sh"]