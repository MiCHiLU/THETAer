# THETAer
An adjuster tool for videos taken by the RICHO THETA S.

initial setup
-------------

First, setup [goenv](https://bitbucket.org/ymotongpoo/goenv). Then:

    $ git clone --branch master --depth 1 https://github.com/MiCHiLU/THETAer.git
    $ mkvirtualenv THETAer
    (THETAer)$ goof make THETAer
    (go:THETAer) (THETAer)$ echo 'goof go go1.6.3\nexport PATH="$GOBIN":"$GOROOT/bin":$PATH' >> activate
    (go:THETAer) (THETAer)$ deactivate && goof workon THETAer
    (go:THETAer) (THETAer)$ make setup

## Get go libraries

    $ go get github.com/samuel/go-gettext

workon
------

    $ workon THETAer
    (THETAer)$ goof workon THETAer
    (go:THETAer) (THETAer)$

build and run
-------------

    $ make && ./bin/theta-er
    theta-er 0.0.0.1 (fd79f0a+, Sep 05 05:21:18 2014, darwin/amd64)
    2014/09/05 13:21:22 Finished: 33.457us

    $ LANG=ja ./bin/theta-er -h
    Usage: ./bin/theta-er [options]

      -i=0: インターバル
      -v=false: デバッグメッセージを表示する

dependencies
------------

* [ffprobe](https://www.ffmpeg.org/ffprobe.html)
* [go-bindata](https://github.com/jteeuwen/go-bindata) [v2.0.4](https://github.com/jteeuwen/go-bindata/releases/tag/2.0.4)
* [pybabel](http://babel.pocoo.org/en/latest/cmdline.html)
