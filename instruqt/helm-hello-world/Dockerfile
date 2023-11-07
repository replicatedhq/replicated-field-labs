FROM gcr.io/kots-field-labs/shell

ARG USERNAME=replicant
ARG USER_UID=1001
ARG USER_GID=$USER_UID

COPY vendor/manifests /home/$USERNAME/manifests
WORKDIR /home/$USERNAME
RUN chmod -R 755 /home/$USERNAME/manifests && \
  chown -R $USERNAME /home/$USERNAME/manifests
