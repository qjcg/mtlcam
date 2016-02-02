#!/usr/bin/env bash
# Download city camera open data

URL_GEOJSON="http://ville.montreal.qc.ca/circulation/sites/ville.montreal.qc.ca.circulation/files/cameras-de-circulation.json"

base_dir=${1:-images}
full_dir="${base_dir}/$(date +'%y%m%d/%H%M%S')"
cache="${base_dir}/cameras-de-circulation.json"
cache_valid_minutes=1
max_concurrency=${2:-90}

# cache expired will store a filename if older than $cache_valid_minutes
cache_expired=$(find -not -newermt "-${cache_valid_minutes} min" -name '*.json')

mkdir -p $full_dir
if [[ ! -r ${cache} || ${cache_expired} ]]; then
	curl -s $URL_GEOJSON -o ${cache}
fi

# aria2c needs "dns" in /etc/nsswitch.conf
# Ref: https://bugzilla.redhat.com/show_bug.cgi?id=1140135
jq -r '. .features[] | .properties."url-image-en-direct"' "${cache}" |
	aria2c -j ${max_concurrency} -d ${full_dir} -i -
