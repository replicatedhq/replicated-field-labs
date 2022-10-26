Replicated Labs - Attendee laptop prep
======================================

* **What you will do**:
  * Set up Replicated CLI Tools for fast iteration
* **Who this is for**: Replicated labs are for anyone who works with app code, docker images, k8s yamls, or does field support for multi-prem applications
  * Full Stack / DevOps / Product Engineers
  * Support Engineers
  * Implementation / Field Engineers
* **Prerequisites**:
  * Basic working knowledge of Kubernetes
  * A Linux or Mac machine on which to set up the development environment (see [this issue](https://github.com/replicatedhq/kots-field-labs/issues/7) for windows)
* **Outcomes**:
  * You will build a working understanding of the Replicated CLI tools and a fast development workflow
  * You will be prepared to integrate the Replicated Vendor platform into your existing CI/CD workflow via GitHub actions or your platform of choice

* * *

## Get started

To start, you'll want to clone this repo somewhere. Optionally, you can fork it first (or you can do this later).

```shell script
git clone https://github.com/replicatedhq/kots-field-labs
cd kots-field-labs/labs/
```

### 1. Install CLI

To start, you'll want to install the `replicated` CLI.
You can install with [homebrew](https://brew.sh) or grab the latest Linux or macOS version from [the replicatedhq/replicated releases page](https://github.com/replicatedhq/replicated/releases).

##### Brew

```shell script
brew install replicatedhq/replicated/cli
```

##### Manual

```shell script
curl -s https://api.github.com/repos/replicatedhq/replicated/releases/latest \
           | grep "browser_download_url.*$(uname | tr '[:upper:]' '[:lower:]')_all.tar.gz" \
           | cut -d : -f 2,3 \
           | tr -d \" \
           | cat <( echo -n "url") - \
           | curl -fsSL -K- \
           | tar xvz replicated
```
Then move `./replicated` to somewhere in your `PATH`:


```shell script
sudo mv replicated /usr/local/bin/
```

##### Verifying

You can verify it's installed with `replicated version`:

```text
$ replicated version
```
```json
{
    "version": "0.40.0",
    "git": "80de6aa",
    "buildTime": "2022-04-05T00:40:18Z",
    "go": {
        "version": "go1.17.8",
        "compiler": "gc",
        "os": "darwin",
        "arch": "amd64"
    }
}
```
