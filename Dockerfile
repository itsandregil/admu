FROM kalilinux/kali-rolling

COPY ./admu /bin/admu

RUN apt-get update && \
    apt-get install -y sudo && \
    groupadd admins && \
    touch /etc/sudoers.d/admins && \
    echo "%admins ALL=(root) NOPASSWD: /bin/admu" > /etc/sudoers.d/admins && \
    chmod 600 /etc/passwd

CMD ["/bin/bash"]
