FROM gcr.io/kots-field-labs/shell

ARG USERNAME=replicant
ARG USER_UID=1001
ARG USER_GID=$USER_UID

# Create the user
RUN groupadd --gid $USER_GID $USERNAME && \
  useradd --uid $USER_UID --gid $USER_GID --shell /bin/bash --create-home $USERNAME

COPY vendor/manifests /home/$USERNAME/manifests
WORKDIR /home/$USERNAME
RUN chmod -R 755 /home/$USERNAME/manifests && \
  chown -R $USERNAME /home/$USERNAME/manifests
