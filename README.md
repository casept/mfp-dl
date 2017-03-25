# mfp-dl
[Windows build](https://ci.appveyor.com/api/projects/status/github/casept/mfp-dl?branch=master&svg=true)

## What is this?   
This is a simple utility intended for downloading all tracks off of [musicforprogramming.net](https://musicforprogramming.net/).
Essentially, it just parses the site's RSS feed
, checks if all tracks contained within are already downloaded, and fetches any that are missing.
Additionally it also downloads the ```folder.jpg``` file, which is needed in order for album art to be displayed properly in some music players. I built this to help me get my feet wet with golang, don't expect it to be rigorously maintained.

## Installation 
Make sure you have installed go and set up your $GOPATH, then do:      

```
go get -u -v github.com/casept/mfp-dl
```

## Usage
```
$ mfp-dl -help
Usage of mfp-dl:
  -dir string
        Directory into which the tracks will be downloaded (default "musicforprogramming")
```      
Example:       
```
mfp-dl -dir ~/Music/musicforprogramming
```

### Development    
This package uses the (not yet) official 'dep' tool. You can install it by running `go get -u github.com/golang/dep`.
Make sure to vendor any dependencies outside the standart library before you submit your pull request.

