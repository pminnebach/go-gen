# The Go generator

A project generator for [Go](http://golang.org/).

See the [Golang documentation](https://golang.org/doc/) for more information.

## Why Go generator ?

Golang team defines [guidelines](https://blog.golang.org/organizing-go-code) to organize your code and structure your application.

The generator builds a standard directory hierarchy for your new golang project.

Optionaly you can build README.md, LICENSE and .gitignore files for when you want to push it to a remote repo like Github.

## Usage

Make sure you have your $GOPATH and $GOBIN set. The Makefile depends on this.

Install Go generator

    go get github.com/pminnebach/go-gen
    cd $GOPATH/src/github.com/pminnebach/go-gen

    go get -u github.com/jteeuwen/go-bindata/...            # necessary for packaging the templates into .go source files.
    go-bindata templates/...

    go install

Run generator

    go-gen [global options] command [command options] [arguments...]

    --repository value, --repo value         Name of your repo. (github.com/<username>)
    --name value, -n value, --project value  Name of the project you want to scaffold (default: "myapp")
    --license value, -l value                Add a LICENSE file. (MIT, Apache)
    --readme, -r                             Add a README.md file
    --gitignore, -g                          Add a .gitignore file.
    --help, -h                               show help
    --version, -v                            print the version


This command will generate a minimalist directory hierarchy plus a makefile and a .gitignore file.

<pre>
myapp
├── .gitignore             # ignores bin, pkg and other useless files
├── LICENSE                # a LICENSE file
├── README.md              # simple readme
├── Makefile               # list available targets with 'make'
├── hello                  # hello package folder
│       ├── hello.go       # hello package
│       └── hello_test.go  # hello test
└── main.go                # main entrypoint of the application

</pre>

A word of warning, go-gen overwrites existing files by default.

### Make

Run unit tests with
```
make test

--> testing...
?   	github.com/pminnebach/myapp	[no test files]
=== RUN   TestSayHello
--- PASS: TestSayHello (0.00s)
PASS
ok  	github.com/pminnebach/myapp/hello	0.007s
```

Clean the project and remove the binary from the $GOBIN folder.
```
make clean

--> cleaning...
Clean OK
```

Build your project and make it run it with
```
make run

--> cleaning...
Clean OK
--> installing...
Install OK
--> running application...
Hello, world.
```

Compile sources and build binary
```
make install

--> cleaning...
Clean OK
--> installing...
Install OK
```

## Known issues

* Due to a bug in [urfave/cli #600](https://github.com/urfave/cli/issues/600), parameters that need a value must also be provided with a value. If the empty parameter is followed by a second parameter, that second parameter will be used as value for the first parameter.

## Contributing

If you would like to submit pull requests, please feel free to apply.

## Dependencies

* Golang
* Make
* [go-bindata](https://github.com/jteeuwen/go-bindata)


### Based on
https://github.com/bench/generator-go
