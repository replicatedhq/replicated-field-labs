# Contributing

*To Do* more in-depth contributing guide. For now, **tl;dr** 1. Follow the style guide 2. PRs welcome!

## Style Guide

Most of these are not intended to be hard-and-fast-rules. In general be thoughtful and do whatever is most conducive to learning the material at hand.

### Prefer exact shell commands to written instructions

For example [Use Heredocs for File Creation](#use-heredocs-for-file-creation)

### Use Heredocs for file creation


**No**

> Create `kots-app.yaml`
> 
> ```yaml
> # kots-app.yaml
> apiVersion: kots.io/v1beta1
> kind: Application
> spec: { ... }
> ```

**Yes**

> Create `kots-app.yaml`
> ```bash
> cat <<EOF > manifests/kots-app.yaml
> apiVersion: kots.io/v1beta1
> kind: Application
> spec: { ... }
> EOF
> ```
  
### Use Diff snippets when explaining edits to be made to files  
  
**No**

> Edit the file `web-deployment.yaml` and change replicas to 2
> ```yaml
> apiVersion: apps/v1
> kind: Deployment
> metadata:
>   name: web-deployment
> spec:
>   replicas: 2 # change this!
> ...
> ```

**Yes**

> Edit the file `web-deployment.yaml` and add a replicas field:
> ```diff
> --- a/manifests/web-deployment.yaml
> +++ b/manifests/web-deployment.yaml
> @@ -8,7 +8,7 @@ spec:
>    selector:
>       matchLabels:
>         app: sentry
> -  replicas: 1
> +  replicas: 2
>    template:
>       metadata:
>         labels:
> ...
> ```

  
