Lab 1 Exercise 0: Hello World
=========================================

This exercise is designed to give you a sandbox to ensure you have the basic CLI tools set up and are prepared to proceed 
with the exercises in Lab 1. 
The README and the YAML sources draw from https://github.com/replicatedhq/replicated-starter-kots

### Get started

To start, you'll want to clone this repo somewhere. Optionally, you can fork it first (or you can do this later).

```shell script
git clone git@github.com:replicatedhq/kots-field-labs
cd kots-field-labs/lab1-kots/lab1-e0-hello-world
```

#### Install CLI

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
           | grep "browser_download_url.*$(uname | tr '[:upper:]' '[:lower:]')_amd64.tar.gz" \
           | cut -d : -f 2,3 \
           | tr -d \" \
           | cat <( echo -n "url") - \
           | curl -fsSL -K- \
           | tar xvz replicated
```
Then move `./replicated` to somewhere in your `PATH`:


```shell script
mv replicated /usr/local/bin/
```

##### Verifying

You can verify it's installed with `replicated version`:

```text
$ replicated version
```
```json
{
  "version": "0.31.0",
  "git": "c67210a",
  "buildTime": "2020-09-03T18:31:11Z",
  "go": {
      "version": "go1.14.7",
      "compiler": "gc",
      "os": "darwin",
      "arch": "amd64"
  }
}
```


#### Configure environment

You'll need to set up two environment variables to interact with vendor.replicated.com:

```
export REPLICATED_APP=...
export REPLICATED_API_TOKEN=...
```

`REPLICATED_APP` should be set to the app slug from the Settings page:

<p align="center"><img src="./doc/REPLICATED_APP.png" width=600></img></p>

Next, create an API token from the [Teams and Tokens](https://vendor.replicated.com/team/tokens) page:

<p align="center"><img src="./doc/REPLICATED_API_TOKEN.png" width=600></img></p>

Ensure the token has "Write" access or you'll be unable create new releases. Once you have the values,
set them in your environment.

```
export REPLICATED_APP=...
export REPLICATED_API_TOKEN=...
```

You can ensure this is working with

```
replicated release ls
```

#### Iterating on your release

Once you've made changes to your manifests, lint them with

```
replicated release lint --yaml-dir=manifests
```

You can push a new release to a channel with

```
make release
```

By default the `Unstable` channel will be used. 
The remaining Lab Exercises use a unique channel per exercise.

## Advanced Usage

### Integrating kurl installer yaml

There is a file `kurl-installer.yaml` that can be used to manage [kurl.sh](https://kurl.sh) installer versions for an embedded Kubernetes cluster. 
You can create a new installer release with

```
replicated installer create --auto
```

Again, `Unstable` will be used by default. The `--promote` flag can be used to override this.

```
replicated installer create --auto --promote=Beta
```

### Tools reference

- [replicated vendor cli](https://github.com/replicatedhq/replicated)

### License

MIT
