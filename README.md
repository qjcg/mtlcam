# mtlcam: Download Montreal traffic camera images

Uses [open data from the City of Montreal](http://donnees.ville.montreal.qc.ca/dataset/cameras-observation-routiere).


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
