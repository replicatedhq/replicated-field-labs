
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
