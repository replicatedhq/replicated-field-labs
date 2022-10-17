---
slug: moving-assets-into-place
type: challenge
title: Moving Assets Into Place
teaser: Getting your airgap assets ready for deployment
notes:
- type: text
  contents: Let's get all our airgap assets ready to deploy
tabs:
- title: Jumpbox
  type: terminal
  hostname: jumpbox
  workdir: /home/replicant
difficulty: basic
timelimit: 800
---

## Moving Assets into Place

Recall the three assets we need for an Air Gap installation:

1. A license with the Air Gap entitlement enabled
2. An Air Gap bundle containing the kURL cluster components
3. An Air Gap bundle containing the application components

We've already begun the download of item (2), since it's the largest
one and we needed some time for it to completed. We also saw how your
customer gets access to all three assets from the Replicated download
portal, then grapped their license file. Now we're going to grab the
application bundle using the command line.

#### Downloading the License File

You can quickly download the license file for any customer using the
`replicated` CLI. This can help you troubleshoot issues a customer is 
having. It also helps us here since we can't easily download a file 
from the Customer Download Portal into our lab jumpbox. Let's get the
license file for our test customer "Replicant".

```
replicated customer download-license --customer Replicant
```
