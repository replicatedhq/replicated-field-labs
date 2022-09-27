FROM gcr.io/kots-field-labs/shell

ARG USERNAME=replicant
ARG USER_UID=1001
ARG USER_GID=$USER_UID

# Create the user
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME
    
# ********************************************************
# * Anything else you want to do like clean up goes here *
# ********************************************************

# [Optional] Set the default user. Omit if you want to keep the default as root.
USER $USERNAME
COPY vendor/manifests /home/$USERNAME/manifests
WORKDIR /home/$USERNAME

ENTRYPOINT [ "su", "-", $USERNAME ]
