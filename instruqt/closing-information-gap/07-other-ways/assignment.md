---
slug: other-ways
id: biqyj5kx0upy
type: challenge
title: Other Ways to Offer a Support Bundle
teaser: Explore other ways to distibute your support bundle and keep it up to date
notes:
- type: text
  contents: How else can you distribute your support bundle
tabs:
- title: Shell
  type: terminal
  hostname: shell
- title: Manifest Editor
  type: code
  hostname: shell
  path: /home/replicant
difficulty: basic
timelimit: 600
---

The support bundle command just needs a specification to run, and it doesn't
matter how it locates it. It can load it from a local file, a URI, or (as we
did above) a secret in the cluster. You can provide the specification at a
publicly available URL to make it easy for them to collect a bundle with
changing their cluster.

Using a URI to Collect a Support Bundle
=======================================

The support bundle specification for this lab is publicly available on a
Google Cloud Storage bucket. To collect a support bundle using that spec, run
the following command:

```
kubectl support-bundle https://storage.googleapis.com/shared-lab-assets/closing-the-information-gap/harbor/support-bundle.yaml
```

You'll see the now familiar analyzer screen showing your customer's cluster.
Since they've now resolved the incident, it will the job service running. If
you look closely, you'll also notice an additional analyzer in this version of
the support bundle.

![Results from Running the Support Bundle from a URI](../assets/analyzers-from-the-uri.png)

You can also use this mechanism to release your support bundle independent of
your application. This can allow the support team to evolve the bundle
independent of your development team and without requiring a full system
release. If you choose to distribute your support bundle at it's own URI, you
should still include it in your application. The support bundle spec allows has
a `uri` property that will be checked before using the specification as is. If
the spec can be retrieved, it will replace the spec that references it.

Including the URI in the Bundle You Distribute
==============================================

Let's finish of our work with the Harbor support bundle by adding a `uri` to
the spec. Edit the file `harbor/troubleshoot/support-bundle.yaml` in the
Manifest Editor and add the URI from the command at the top of the `spec`,
before the collectors.

```
  uri: https://storage.googleapis.com/shared-lab-assets/closing-the-information-gap/harbor/support-bundle.yaml
```

Don't forget to the save the file.

![Adding a URI to the Support Bundle spec](../assets/adding-a-uri-to-the-spec.png)

You should also bump the version number for the release to `16.9.1` in the
`Chart.yaml` file, repackage your chart, and delete the old tarball from your
release directory. The following commands will take care of that for you.


```
yq -i '.version = "16.9.1"' harbor/Chart.yaml
helm package harbor --destination ./release
rm ./release/harbor-16.9.0.tgz
```

Now you need to release your update and promote it. Remember that for a real
application you won't be instantly promoting to `Beta` and `Stable`. Your
release sequence will be `5` unless you've played around with additional
releases. If you have, use the current sequence.

```
replicated release create --promote Unstable --yaml-dir ./release --version 16.9.1  \
  --release-notes "Provides the ability to upgrade the support bundle independent of the application"
replicated release promote 5 Beta --version 16.9.1 \
  --release-notes "Provides the ability to upgrade the support bundle independent of the application"
replicated release promote 5 Stable --version 16.9.1 \
  --release-notes "Provides the ability to upgrade the support bundle independent of the application"
```

Running the New Bundle as Geeglo
================================

You can now resume acting as Geeglo to upgrade your installation and collect a new support bundle containing your

```
kubectl support-bundle --load-cluster-specs
```

Note that you didn't add a new analyzer to the support bundle defintion in your
Helm chart, but you still see the new analyzer in your results this is because
the `support-bundle` command checked the URI for an upgrade, found one, and
used in instead of the out-of-date specification.

![Additional Analzer Result After Spec is Updated](../assets/additional-analyzer-result.png)

