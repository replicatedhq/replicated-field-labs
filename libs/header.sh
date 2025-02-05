
show_credentials () {
    CYAN='\033[0;36m'
    GREEN='\033[1;32m'
    NC='\033[0m' # No Color
    password=$(get_password)
    echo -e "${GREEN}Credentials for ${CYAN}https://vendor.replicated.com"
    echo -e "${GREEN}Username: ${CYAN}${INSTRUQT_PARTICIPANT_ID}@replicated-labs.com"
    echo -e "${GREEN}Password: ${CYAN}${password}${NC}"
}

get_replicated_sdk_version () {
  set +eu pipefail
  replicated_sdk_version=$(agent variable get REPLICATED_SDK_VERSION)

  # if we don't already have a token, fetch one
  if [[ -z "$replicated_sdk_version" ]]; then
    set -eu pipefail
    token=$(curl --silent "https://registry.replicated.com/v2/token?scope=repository:library/replicated:pull&service=registry.replicated.com" | jq -r .token)
    replicated_sdk_version=$(curl --silent -H "Authorization: Bearer ${token}" https://registry.replicated.com/v2/library/replicated/tags/list | jq -r '.tags[]' | awk -F '[.-]' '{
        # Extract version components
        major=$1;
        minor=$2;
        patch=$3;
        prerelease=$4;
        prerelease_number=$5;

        # Assign priority to pre-release versions
        if (prerelease == "alpha") {
          prerelease_priority = 1;
        } else if (prerelease == "beta") {
          prerelease_priority = 2;
        } else {
          prerelease_priority = 3;
        }

        # Handle missing pre-release number
        if (prerelease_number == "") {
          prerelease_number = 0;
        }

        # Format output to aid sorting
        printf "%04d%04d%04d%02d%04d-%s\n", major, minor, patch, prerelease_priority, prerelease_number, $0
      }' | sort -r | head -1 | sed 's/^[0-9]*-//')
  fi

  set -eu
  echo ${replicated_sdk_version}
}

get_embedded_cluster_version () {
  set +eu pipefail
  embedded_cluster_version=$(agent variable get EMBEDDED_CLUSTER_VERSION)

  # if we don't already have a token, fetch one
  if [[ -z "$empedded_cluster_version" ]]; then
    embedded_cluster_version=$(curl -s "https://api.github.com/repos/replicatedhq/embedded-cluster/releases/latest" | jq -r .tag_name)
  fi

  set -eu pipefail
  echo ${embedded_cluster_version}
}

get_username () {
  echo ${INSTRUQT_PARTICIPANT_ID}@replicated-labs.com
}

get_password () {
    password=$(echo -n "${INSTRUQT_PARTICIPANT_ID}" | sha256sum)
    echo ${password::20}
}

get_api_token () {
  set +eu
  access_token=$(agent variable get REPLICATED_API_TOKEN)

  # if we don't already have a token, fetch one
  if [[ -z "$access_token" ]]; then
    set -eu
    sleep 5
    password=$(get_password)
    login=$( jq -n -c --arg email "${INSTRUQT_PARTICIPANT_ID}@replicated-labs.com" --arg password "${password}" '$ARGS.named' )
    set +eu pipefail
    token=$(curl -s -H "Content-Type: application/json" --request POST -d "$login" https://api.replicated.com/vendor/v1/login | jq -r ".token")
    set -eu pipefail
    
    i=0
    while [[ ( -z "$token" || "$token" == "null" ) && $i -lt 20 ]]
    do
        sleep $((i*5))
        set +eu pipefail
        token=$(curl -s -H "Content-Type: application/json" --request POST -d "$login" https://api.replicated.com/vendor/v1/login | jq -r ".token")
        set -eu pipefail
        i=$((i+1))
    done

    UUID=$(cat /proc/sys/kernel/random/uuid)
    apiToken=$( jq -n -c --arg name "instruqt-${UUID}" --argjson read_only false '$ARGS.named' )
    access_token=$(curl -s -H "Content-Type: application/json" -H "Authorization: $token" --request POST -d "$apiToken" https://api.replicated.com/vendor/v1/user/token | jq -r ".access_token")

    agent variable set REPLICATED_API_TOKEN $access_token
  fi
  set +eu
  echo ${access_token}
}

get_app_slug () {
  application=${1:-"Slackernews"}
  access_token=$(get_api_token)
  app_slug=$(curl --header 'Accept: application/json' --header "Authorization: ${access_token}" https://api.replicated.com/vendor/v3/apps | jq -r --arg application ${application} '.apps[] | select( .name | startswith( $application )) | .slug')
  echo ${app_slug}
}

get_app_id () {
  application=${1:-"Slackernews"}
  access_token=$(get_api_token)
  app_id=$(curl --header 'Accept: application/json' --header "Authorization: ${access_token}" https://api.replicated.com/vendor/v3/apps | jq -r --arg application ${application} '.apps[] | select( .name | startswith( $application )) | .id')
  echo ${app_id}
}

get_customer_id () {
  customer=${1}
  access_token=$(get_api_token)
  app_id=$(get_app_id)
  customer_id=$(curl --header 'Accept: application/json' --header "Authorization: ${access_token}" https://api.replicated.com/vendor/v3/app/${app_id}/customers | jq -r --arg name $customer '.customers[] | select ( .name == $name ) | .id')
  echo ${customer_id}
}

get_license_id () {
  customer=${1}
  access_token=$(get_api_token)
  app_id=$(get_app_id)
  license_id=$(curl --header 'Accept: application/json' --header "Authorization: ${access_token}" https://api.replicated.com/vendor/v3/app/${app_id}/customers | jq -r --arg name $customer '.customers[] | select ( .name == $name ) | .installationId')
  echo ${license_id}
}

get_admin_console_password() {
  password=$(echo -n "${INSTRUQT_PARTICIPANT_ID}:${INSTUQT_CHALLENGE_ID}" | sha256sum)
  echo ${password::20}
}

get_slackernews_domain() {
  echo cluster-30443-${INSTRUQT_PARTICIPANT_ID}.env.play.instruqt.com
}

get_slackernews() {
  # get the app slug, since there's only one app created by the automation, just grab the first in the list
  app_slug=$(get_app_slug)

  # grab the sources for the Helm chart using a community license
  helm registry login chart.slackernews.io --username marc@replicated.com --password 2ViYIi8SDFubA8XwQRhJtcrwn4C
  helm pull --untar oci://chart.slackernews.io/slackernews/slackernews

  # specify the nodeport for NGINX so we get a consistent and addressable endpoint
  # TODO: Update upstream to take this as a value
  sed -i '17 a\    nodePort: 30443' slackernews/templates/nginx-service.yaml
 
  # remove the Replicated SDK dependency, if we add more dependencies to
  # Slackernews this will need to be revised
  yq -i 'del(.dependencies)' slackernews/Chart.yaml

  # start version numbers over to simplify the lab text
  yq -i '.version = "0.1.0"' slackernews/Chart.yaml

  # get rid of troubleshoot files since leaners will create their own
  rm -rf slackernews/troubleshoot slackernews/templates/preflights.yaml slackernews/templates/support-bundle.yaml

  # set the values file ot use the right proxy image URI
  web_image=$(yq .images.slackernews.repository slackernews/values.yaml)
  rewritten_web_image=${web_image//images.slackernews.io/proxy.replicated.com}
  rewritten_web_image=${rewritten_web_image//proxy\/slackernews/proxy\/${app_slug}}
  yq -i ".images.slackernews.repository = \"${rewritten_web_image}\"" slackernews/values.yaml

  nginx_image=$(yq .images.nginx.repository slackernews/values.yaml)
  rewritten_nginx_image=${nginx_image//images.slackernews.io/proxy.replicated.com}
  rewritten_nginx_image=${rewritten_nginx_image//proxy\/slackernews/proxy\/${app_slug}}
  yq -i ".images.nginx.repository = \"${rewritten_nginx_image}\"" slackernews/values.yaml

  # add some optional components to make the application a bit more representative
  yq -i '.nginx.enabled = true' slackernews/values.yaml
  yq -i '.postgres.deploy_postgres = true' slackernews/values.yaml
  yq -i '.postgres.enabled = true' slackernews/values.yaml
  yq -i '.postgres.password = "thisisasecret"' slackernews/values.yaml

  # address awkward scenario where a TLS cert is required even if TLS isn't enabled
  # TODO: Fix upstream to not require TLS certs uneless TLS is enabled
  openssl req -new -newkey rsa:4096 -days 365 -nodes -x509 -keyout server.key -out server.crt -subj "/CN=Slackernews" -addext "subjectAltName = DNS:$(get_slackernews_domain)" \
    && yq -i ".service.tls.key = \"$(cat server.key)\"" slackernews/values.yaml \
    && rm server.key \
    && yq -i ".service.tls.cert = \"$(cat server.crt)\"" slackernews/values.yaml \
    && rm server.crt
 
  # since we have the certs anyway, let's enable TLS
  yq -i '.service.tls.enabled = true' slackernews/values.yaml

  # let's also deelte the values injected by Replicated so users can release
  # the chart without any sort of double injection
  yq -i 'del(.replicated)' slackernews/values.yaml
  yq -i 'del(.global.replicated)' slackernews/values.yaml
}
