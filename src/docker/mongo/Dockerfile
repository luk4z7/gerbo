FROM mongo

RUN apt-get update && apt-get install -y telnet

# Export xterm
RUN echo export TERM=xterm >> ~/.bashrc
RUN apt-get install -y cron vim
ADD remove.sh /
RUN chmod +x /remove.sh