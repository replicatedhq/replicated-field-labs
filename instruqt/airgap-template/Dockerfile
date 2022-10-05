FROM gcr.io/kots-field-labs/shell
# First we update the apt package manager cache

RUN apt-get update && apt-get -y install curl apt-transport-https ca-certificates gnupg

# Use the gcloud installation process described in the docs:
# https://cloud.google.com/sdk/docs/install#deb
RUN echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && \
  curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg add -

# Update apt again and install gcloud, kubectl and helm
RUN apt -y update && \
  apt -y install google-cloud-sdk && \
  apt -y clean


