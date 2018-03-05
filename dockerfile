FROM debian:sid-slim

RUN cd && set -ex \
  echo 'Acquire::http {No-Cache=True;};' | tee /etc/apt/apt.conf.d/no-http-cache && \
  apt-get update && \
  apt-get -y install --no-install-recommends golang golang-go.tools \
    bash-completion curl git vim vim-gocomplete apt-transport-https ca-certificates debootstrap && \
  curl -Ls https://github.com/kura/go-bash-completion/raw/master/etc/bash_completion.d/go -o /etc/bash_completion.d/go && \
  ln -sf /etc/skel/.bashrc /root/.bashrc && \
  apt-get autoremove -y && apt-get clean && \
  rm -rf -- /tmp/* /var/tmp/* \
     /var/cache/apt/* /var/lib/apt/lists/*

CMD bash
