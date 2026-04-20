FROM kalilinux/kali-rolling

COPY ./admu /bin/admu

RUN apt-get update && apt-get upgrade -y
RUN apt-get install -y sudo
RUN groupadd admins
RUN touch /etc/sudoers.d/admins
RUN echo "%admins ALL=(root) NOPASSWD: /bin/admu" > /etc/sudoers.d/admins

CMD ["/bin/bash"]
