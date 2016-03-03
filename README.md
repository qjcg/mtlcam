# mtlcam

Pull down the latest Montreal traffic camera images using the [open GeoJSON
data](http://donnees.ville.montreal.qc.ca/dataset/cameras-observation-routiere)
provided by the City of Montreal.


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
  -d    print debug messages
  -p string
        parent directory for downloaded files (default "images")
  -v    print version

$ mtlcam
[...]

$ ls -R images
[...]
```


## License

MIT.
