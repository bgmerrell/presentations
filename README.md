presentations
=============

Miscellaneous presentations

How to view
-----------
To view these slide shows, you need to make sure that you have installed:
- git
- go

#### Linux
- Archlinux: `sudo pacman -Sy git go`
- RHEL/Fedora: `sudo yum install git go`
- Debian: `sudo apt-get install git go`

#### Mac OS X
- Git: http://git-scm.com/download/mac
- Go: https://code.google.com/p/go/wiki/Downloads?tm=2

After you have made sure that both of these are installed, and that you
have [set your GOPATH][1], you can then get the presentation server
running after you've executed the following commands:
```bash
go get code.google.com/p/go.tools/cmd/present
git clone https://github.com/bgmerrell/presentations
cd presentations
$GOPATH/bin/present
```


[1]: http://golang.org/doc/code.html#GOPATH "Golang's website has some excellent documentation on how to do this"
