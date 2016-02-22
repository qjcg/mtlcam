# mtlcam: Download Montreal traffic camera images

## Install

```
go get -u github.com/qjcg/mtlcam
```

## Usage

```
$ mtlcam -h

mtlcam: Download Montreal traffic camera images
Data source: http://donnees.ville.montreal.qc.ca/dataset/cameras-observation-routiere

  -c int
        max concurrent downloads (default 90)
  -d string
        parent directory for downloaded files (default "images")

$ mtlcam
[...]

$ ls -R images
[...]
```

## License

MIT.
