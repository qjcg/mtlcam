#!/usr/bin/env bash
# Download city camera open data

URL_GEOJSON="http://ville.montreal.qc.ca/circulation/sites/ville.montreal.qc.ca.circulation/files/cameras-de-circulation.json"

base_dir=${1:-images}
full_dir="${base_dir}/$(date +'%y%m%d/%H%M%S')"
cache="${base_dir}/cameras-de-circulation.json"
max_concurrency=${2:-10}

mkdir -p $full_dir
if [[ ! -r ${cache} ]]; then
	curl -s $URL_GEOJSON -o ${cache}
fi

cat "${cache}" |
	jq '. .features[] | .properties."url-image-en-direct"' |
	xargs -n ${max_concurrency} wget -P ${full_dir} -i -
