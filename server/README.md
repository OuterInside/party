Party Server
============

How to build
------------
### symbolic link to $GOPATH
```
GO_PARTY_PKG=$(echo $GOPATH | cut -d: -f1)/src/github.com/OuterInside/party
mkdir -p $GO_PARTY_PKG
ln -s `pwd` $GO_PARTY_PKG
cd $GO_PARTY_PKG/server
```

### install vendoring tool (glide)
```
brew install glide
```

ref: https://github.com/Masterminds/glide#install

### build
```
glide install
go build -o server .
```

### run
```
./server -port 8080
```
