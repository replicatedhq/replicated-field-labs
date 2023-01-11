#!/usr/bin/env bash

ARGS=$(getopt -o elsvar --long storage,service,limits,entrypoint,admin-console,registry  -- "$@")

get_deployment() {
# get a random deployment name from the application deployed by a vendor
kubectl get deployments -l kots.io/app-slug=annarchy  | awk '{if(NR>1)print $1}' | xargs shuf -n1 -e
}


if [[ $? -ne 0 ]]; then
  echo "Failed parsing options." >&2
  exit 1
fi

eval set -- "$ARGS"

while [ : ]; do
# do some chaos

  case "${1}" in
    -e | --entrypoint)
    # patch container's entrypoint so it fails to start
      shift
      ;;
    -l | --limits)
    # patch resource limits to be very low
      deployment=$(get_deployment)
      kubectl patch deployment "${deployment}" --type='json' -p='[{"op": "replace", "path": "/spec/template/spec/containers/0/resources/requests/memory", "value": "10Mi"}]'
      kubectl patch deployment "${deployment}" --type='json' -p='[{"op": "replace", "path": "/spec/template/spec/containers/0/resources/limits/memory", "value": "10Mi"}]'
      shift
      ;;
    -s | --storage)
    # patch StorageClass so PVCs fail
      shift
      ;;
    -v | --service)
    # patch Services so lb connectivity fails
      shift
      ;;
    -a | --admin-console)
    # break Admin Console so GUI doesn't work
      shift
      ;;
    -r | --registry)
    # break registry so images can't be pulled
      shift
      ;;
    --)
      shift
      break
      ;;
  esac
done











