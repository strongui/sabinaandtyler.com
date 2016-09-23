#!/usr/bin/env bash
set -eu

base_url="http://mollyandnate.com/"
folder="mollyandnate.com"

if [ "${BASE_URL:-}" != "" ] ; then
  base_url=${BASE_URL:-}
fi

if [ "${FOLDER:-}" != "" ] ; then
  folder=${FOLDER:-}
fi

echo "BASE_URL: $base_url"
echo "FOLDER: $folder"

rm -rf "$folder"
mkdir "$folder"

hugo --source="site" \
     --baseURL="$base_url" \
     --destination="../$folder" \
     --disableRSS=true \
     --disableSitemap=true

rm -rf "$folder/index/"
