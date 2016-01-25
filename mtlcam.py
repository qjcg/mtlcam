#!/usr/bin/env python3
# Download Montreal city traffic camera open data

import argparse
import concurrent.futures
import datetime
import json
import os

import requests

URL_GEOJSON = "http://ville.montreal.qc.ca/circulation/sites/ville.montreal.qc.ca.circulation/files/cameras-de-circulation.json"

# format string passed to datetime.strftime
FMT_DATE = "%y%m%d"
FMT_TIME = "%H%M%S"

CACHE_FILE = "cameras-de-circulation.json"


def get_urls(json_data):
    return [url for url in (site['properties']['url-image-en-direct'] for site in json_data['features'])]


def download_image(urls):
    futures = []
    for url in urls:
        with concurrent.futures.ThreadPoolExecutor(max_workers=10) as executor:
            print("GET {}".format(url))
            future = executor.submit(requests.get, url)
            futures.append(future)

            img_filename = os.path.basename(img_url)
            with open(os.path.join(fulldir, img_filename), 'bw') as f:
                f.write(res.content)


# TODO: Only update cache if time delta larger than some value
def get_data(base_dir):
    """Update GeoJSON cache."""
    res = requests.get(URL_GEOJSON)
    data = res.json()
    with open(os.path.join(base_dir, "cache.json"), 'w') as f:
        json.dump(data, f)

    return data


def make_dirs(base_dir):
    """Create timestamped directory for downloading images."""
    now = datetime.datetime.now()
    datestamp = now.strftime(FMT_DATE)
    timestamp = now.strftime(FMT_TIME)
    fulldir = os.path.join(base_dir, datestamp, timestamp)
    os.makedirs(fulldir, exist_ok=True)

    return fulldir


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument("-d", "--dir", default="mtlcam", help="base dir for downloading images")
    parser.add_argument("-c", "--concurrency", default=10, type=int, help="max concurrency")
    conf = parser.parse_args()

    fulldir = make_dirs(conf.dir)
    data = get_data(conf.dir)
    urls = get_urls(data)
    download_images(urls)
