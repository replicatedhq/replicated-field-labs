#!/usr/bin/env bash

# get a random deployment name from the application deployed by a vendor

deployment=$(kubectl get deployments -l kots.io/app-slug=annarchy  | awk '{if(NR>1)print $1}' | xargs shuf -n1 -e)

kubectl patch deployment $deployment --type='json' -p='[{"op": "replace", "path": "/spec/template/spec/containers/0/resources/requests/memory", "value": "10Mi"}]'
