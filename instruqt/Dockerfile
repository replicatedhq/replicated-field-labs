FROM replicated/vendor-cli AS vendor-cli

# Troubleshoot plugins
FROM ubuntu:jammy AS krew

RUN apt-get update && apt-get install -y ca-certificates curl git

RUN curl https://krew.sh/preflight | bash
RUN curl https://krew.sh/support-bundle | bash

FROM ubuntu:jammy 

RUN apt-get update && apt-get install -y \
  apt-transport-https \
  gnupg \
  ca-certificates \
  curl \
  git \
  jq

# kubectl
RUN curl -fsSL https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - && \
  echo "deb [arch=amd64] https://apt.kubernetes.io/ kubernetes-xenial main" | tee /etc/apt/sources.list.d/kubernetes.list
RUN apt-get update && apt-get install -y kubectl 

RUN apt-get clean

# Helm
RUN curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash

# YQ, using the github release
RUN curl -s https://api.github.com/repos/mikefarah/yq/releases/latest | \
  jq '.assets[] | select(.name == "yq_linux_amd64").browser_download_url' | \
  xargs curl --location --output /usr/local/bin/yq && \
  chmod 755 /usr/local/bin/yq

# Troubleshoot plugins
COPY --from=krew /root/.krew/bin/kubectl-preflight /usr/local/bin
COPY --from=krew /root/.krew/bin/kubectl-support_bundle /usr/local/bin

# sbctl
RUN curl -s https://api.github.com/repos/replicatedhq/sbctl/releases/latest | \
  jq '.assets[] | select(.name == "sbctl_linux_amd64.tar.gz").browser_download_url' | \
  xargs curl --location --output /usr/local/bin/sbctl && \
  chmod 755 /usr/local/bin/sbctl


# replicated CLI
COPY --from=vendor-cli /replicated /usr/local/bin

