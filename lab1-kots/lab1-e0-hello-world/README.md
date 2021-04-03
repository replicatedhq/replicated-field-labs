Lab 1 Exercise 0: Hello World
=========================================

This exercise is designed to give you a sandbox to ensure you have the basic CLI tools set up and are prepared to proceed 
with the exercises in Lab 1. 
The README and the YAML sources draw from https://github.com/replicatedhq/replicated-starter-kots

### Get started

This repo is a [GitHub Template Repository](https://help.github.com/en/articles/creating-a-repository-from-a-template). You can create a private copy by using the "Use this Template" link in the repo:

![Template Repo](https://help.github.com/assets/images/help/repository/use-this-template-button.png)

You should use the template to create a new **private** repo in your org, for example `mycompany/kots-app` or `mycompany/replicated-starter-kots`.

Once you've created a repository from the template, you'll want to `git clone` your new repo and `cd` into it locally.


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
replicated release create --auto
```

By default the `Unstable` channel will be used. You can override this with the `--promote` flag:

```
replicated release create --auto --promote=Beta
```


### Integrating with CI

This repo contains a [GitHub Actions](https://help.github.com/en/github/automating-your-workflow-with-github-actions/about-github-actions) workflow for ci at [./.github/workflows/main.yml](./.github/workflows/main.yml). You'll need to [configure secrets](https://help.github.com/en/github/automating-your-workflow-with-github-actions/virtual-environments-for-github-actions#creating-and-using-secrets-encrypted-variables) for `REPLICATED_APP` and `REPLICATED_API_TOKEN`. On every push this will:

- Ensure a channel exists for the branch that was pushed to
- Create a release based on the contents of `./manifests`

## Advanced Usage

### Integrating kurl installer yaml

There is a file `kurl-installer.yaml` that can be used to manage [kurl.sh](https://kurl.sh) installer versions for an embedded Kubernetes cluster. This will be automatically released in CI. You can create a release manually with

```
replicated installer create --auto
```

### Tools reference

- [replicated vendor cli](https://github.com/replicatedhq/replicated)

### License

MIT
