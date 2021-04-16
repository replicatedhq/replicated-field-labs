Single Player Mode
======================


This guide will walk you through running all the labs on your own

## 1. Prerequisites

* Create a vendor.replicated.com account
* Create and copy an API token, set it in your shell with `export REPLICATED_API_TOKEN=...`
* Choose a unique name for your session, e.g. `dh-test`

## 2. Create an environment JSON

Copy the example file and edit it with your name and slug.
Since single-player mode requires you to have server access to provision the boxes in the first place,
you can

```
cp environments_test.json env-dh-test.json
vim env-dh-test.json
```
