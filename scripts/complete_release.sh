#! /bin/bash

set -e

echo "Getting latest release..."


RELEASE_RES=$(curl -H "Authorization: token $GITHUB_TOKEN" -s "https://api.github.com/repos/Octy-ai/octy-cli/releases/latest")
RELEASE_ID=$( jq -r  '.id' <<< "${RELEASE_RES}" ) 
RELEASE_BODY=$( jq -r  '.body' <<< "${RELEASE_RES}" | sed -e 's/^"//' -e 's/"$//')


if [[ "${RELEASE_ID}" == null ]]; then
    printf "No release found! Check all required ENV variables have been set and that the GITHUB_TOKEN has not expired!" >&2
    exit 1 
fi

UPDATED_BODY="${RELEASE_BODY} (Assets Updated)"

echo "Updating latest release with ID: $RELEASE_ID"
echo "Updating latest release body to : $UPDATED_BODY"

generate_post_data()
{
  cat <<EOF
{
  "body":"$UPDATED_BODY"
}
EOF
}

curl \
  -X PATCH \
  -H "Accept: application/vnd.github.v3+json" \
  -H "Authorization: token $GITHUB_TOKEN" \
  https://api.github.com/repos/Octy-ai/octy-cli/releases/$RELEASE_ID \
  -d "$(generate_post_data)"

echo "Updated release with ID: $RELEASE_ID"