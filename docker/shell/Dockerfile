FROM replicated/vendor-cli AS vendor-cli

# Troubleshoot plugins
FROM ubuntu:jammy AS installer

RUN apt-get update && apt-get install -y \
  apt-transport-https \
  gnupg \
  ca-certificates \
  curl \
  git \
  jq

RUN apt-get update && apt-get install -y ca-certificates curl git jq

FROM installer AS krew
RUN curl https://krew.sh/preflight | bash
RUN curl https://krew.sh/support-bundle | bash

# sbctl
FROM installer AS sbctl
RUN apt-get update && apt-get install -y ca-certificates curl jq
RUN curl -s https://api.github.com/repos/replicatedhq/sbctl/releases/latest | \
  jq '.assets[] | select(.name == "sbctl_linux_amd64.tar.gz").browser_download_url' | \
  xargs curl --location | \
  tar -C /usr/local/bin -xzvf - sbctl && \
  chmod 755 /usr/local/bin/sbctl

# kubctl
FROM installer AS kubectl

# kubectl
# If the folder `/etc/apt/keyrings` does not exist, it should be created before the curl command, read the note below.
RUN curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.30/deb/Release.key | gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg && \
  chmod 644 /etc/apt/keyrings/kubernetes-apt-keyring.gpg 

RUN echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.30/deb/ /' | \
  tee /etc/apt/sources.list.d/kubernetes.list && chmod 644 /etc/apt/sources.list.d/kubernetes.list
RUN apt-get update && apt-get install -y kubectl 

# final image
FROM ubuntu:jammy


ARG USERNAME=replicant
ARG USER_UID=1001
ARG USER_GID=$USER_UID

# Create the user
RUN groupadd --gid $USER_GID $USERNAME && \
  useradd --uid $USER_UID --gid $USER_GID --shell /bin/bash --create-home $USERNAME

RUN apt-get update && apt-get install -y \
  ca-certificates \
  curl \
  git \
  jq \
  tmux

# add kubectl 
COPY --from=kubectl /usr/bin/kubectl /usr/local/bin

# Helm
RUN curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash

# YQ, using the github release
RUN curl -s https://api.github.com/repos/mikefarah/yq/releases/latest | \
  jq '.assets[] | select(.name == "yq_linux_amd64").browser_download_url' | \
  xargs curl --location --output /usr/local/bin/yq && \
  chmod 755 /usr/local/bin/yq

# add  KOTS
RUN curl -fsSL https://kots.io/install | bash

# add Troubleshoot plugins
COPY --from=krew /root/.krew/bin/kubectl-preflight /usr/local/bin
COPY --from=krew /root/.krew/bin/kubectl-support_bundle /usr/local/bin

# add sbctl
COPY --from=sbctl /usr/local/bin/sbctl /usr/local/bin

# add replicated CLI
COPY --from=vendor-cli /replicated /usr/local/bin

