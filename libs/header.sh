
show_credentials () {
    CYAN='\033[0;36m'
    GREEN='\033[1;32m'
    NC='\033[0m' # No Color
    password=$(get_password)
    echo -e "${GREEN}Credentials for ${CYAN}https://vendor.replicated.com"
    echo -e "${GREEN}Username: ${CYAN}${INSTRUQT_PARTICIPANT_ID}@replicated-labs.com"
    echo -e "${GREEN}Password: ${CYAN}${password}${NC}"
}

get_username () {
  echo ${INSTRUQT_PARTICIPANT_ID}@replicated-labs.com
}

get_password () {
    password=$(echo -n "${INSTRUQT_PARTICIPANT_ID}" | sha256sum)
    echo ${password::20}
}

get_api_token () {
  password=$(get_password)
  login=$( jq -n -c --arg email "${INSTRUQT_PARTICIPANT_ID}@replicated-labs.com" --arg password "${password}" '$ARGS.named' )
  set +e pipefail
  token=$(curl -s -H "Content-Type: application/json" --request POST -d "$login" https://id.replicated.com/v1/login | jq -r ".token")
  set -e pipefail
  

  i=0
  while [[ "$token" == "null" && $i -lt 20 ]]
  do
      sleep 2
      set +e pipefail
      token=$(curl -s -H "Content-Type: application/json" --request POST -d "$login" https://id.replicated.com/v1/login | jq -r ".token")
      set -e pipefail
      i=$((i+1))
  done

  UUID=$(cat /proc/sys/kernel/random/uuid)
  apiToken=$( jq -n -c --arg name "instruqt-${UUID}" --argjson read_only false '$ARGS.named' )
  access_token=$(curl -s -H "Content-Type: application/json" -H "Authorization: $token" --request POST -d "$apiToken" https://api.replicated.com/vendor/v1/user/token | jq -r ".access_token")

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
