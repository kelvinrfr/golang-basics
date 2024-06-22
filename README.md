# golang-basics


```
go build # transforms script in a executable file
go test # run test files that end with "_test"
```


discover the install path by running the go list command https://go.dev/cmd/go/#hdr-List_packages_or_modules
```
 go list -f '{{.Target}}'
```

On Linux or Mac, run the following command:
```
$ export PATH=$PATH:/path/to/your/install/directory
```

On Windows, run the following command:
```
$ set PATH=%PATH%;C:\path\to\your\install\directory
```

Once you've updated the shell path, run the go install command to compile and install the package.
```
$ go install
```