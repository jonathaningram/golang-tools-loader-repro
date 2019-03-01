# Repro when using Go modules for potential discrepancy between go generate and running the generate command directly

## What env am I using?

```
$ go version
go version go1.12 darwin/amd64
```

```
$ go env
GOARCH="amd64"
GOBIN=""
GOCACHE="/Users/Me/Library/Caches/go-build"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOOS="darwin"
GOPATH="/Users/Me/go"
GOPROXY=""
GORACE=""
GOROOT="/usr/local/go"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/darwin_amd64"
GCCGO="gccgo"
CC="clang"
CXX="clang++"
CGO_ENABLED="1"
GOMOD="/Users/Me/code/golang-tools-loader-repro/go.mod"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/rf/9_n6qmkn13bbq6486y2yf6y00000gq/T/go-build057434341=/tmp/go-build -gno-record-gcc-switches -fno-common"
```

## Set up

Check out the code outside of `$GOPATH`, e.g. at `/Users/Me/code/repro`.

```
cd cmd/greetings-gen
go install # installs into $GOPATH/bin
which greetings-gen # just check it's there
```

## Generate does not work

_Note:_ before running make sure that `github.com/twitchtv/twirp` does not exist in your `$GOPATH` (Twirp is used as some arbitrary dependency that this repro's package needs—I don't believe that Twirp has anything to do with this):

```
$ ls $GOPATH/src/github.com/twitchtv
<nothing>
```

Now run the generator:

```
cd <repro-root>, e.g. /Users/Me/code/repro
greetings-gen -pkg x.com/x/greetings
```

Output is bad:

```
/Users/Me/go/pkg/mod/github.com/twitchtv/twirp@v5.5.2+incompatible/context.go:21:2: could not import github.com/twitchtv/twirp/internal/contextkeys (cannot find package "github.com/twitchtv/twirp/internal/contextkeys" in any of:
	/usr/local/go/src/github.com/twitchtv/twirp/internal/contextkeys (from $GOROOT)
	/Users/Me/go/src/github.com/twitchtv/twirp/internal/contextkeys (from $GOPATH))
2019/03/01 17:15:20 can't generate: can't load program from package "x.com/x/greetings": couldn't load packages due to errors: github.com/twitchtv/twirp
```
