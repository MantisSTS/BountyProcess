FROM debian:latest

RUN apt update && apt upgrade -y
RUN apt install bash wget htop -y
RUN apt install -y python3 python3-pip
RUN pip install amqp
RUN pip install pika
RUN wget https://go.dev/dl/go1.20.linux-amd64.tar.gz -O /tmp/go.tar.gz && \
    tar -C /usr/local -xzf /tmp/go.tar.gz && \
    rm /tmp/go.tar.gz
ENV PATH="${PATH}:/usr/local/go/bin/"
ENV PATH="${PATH}:/root/go/bin/"
RUN apt install -y git
RUN mkdir /opt/tools
RUN git clone https://github.com/EnableSecurity/wafw00f /opt/tools/wafw00f && cd /opt/tools/wafw00f && python3 setup.py install
RUN /usr/local/go/bin/go install -v github.com/projectdiscovery/nuclei/v2/cmd/nuclei@latest
RUN /root/go/bin/nuclei -update-templates
RUN /usr/local/go/bin/go install -v github.com/projectdiscovery/httpx/cmd/httpx@latest
RUN /usr/local/go/bin/go install -v github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest
RUN /usr/local/go/bin/go install -v github.com/projectdiscovery/chaos-client/cmd/chaos@latest
RUN /usr/local/go/bin/go install -v github.com/OWASP/Amass/v3/...@master
RUN /usr/local/go/bin/go install -v github.com/tomnomnom/assetfinder@latest
RUN git clone https://github.com/aboul3la/Sublist3r.git /opt/tools/Sublist3r && cd /opt/tools/Sublist3r && pip3 install -r ./requirements.txt && python3 setup.py install


COPY subdomain /root/app/workers/subdomain

WORKDIR /root/app/workers/subdomain

CMD ["python3 /root/app/workers/subdomain/subdomain_worker.py"]