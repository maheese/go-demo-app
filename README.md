# go-demo-app
![CodeQL](https://github.com/maheese/go-demo-app/workflows/CodeQL/badge.svg) ![Build, Test, Deploy](https://github.com/maheese/go-demo-app/workflows/Build,%20Test,%20Deploy/badge.svg)

This is a simple go webapp for demonstrating how to use GitHub Actions to deploy to Cloud.gov. 

#### Dependencies

- [git][1]
- [go][2]
- [cf-cli][3]

#### Clone the repository

Clone the repository and `cd` into it:

```shell
git clone https://github.com/maheese/go-demo-app
cd go-demo-app
```

#### Testing the app
You can run the dummy unit tests.

```shell
$ go test ''
ok      /go-demo-app       0.413s
```

#### Running the app locally
You can access the app at `http://localhost:8080`

```shell
$ go build
$ ./go-demo-app.exe
Listening on port 8080
```

#### Pushing the app to cloud.gov locally

```shell
$ cf push .
Pushing app . to org / space ...
Applying manifest manifest.yml...
Manifest applied
Packaging files to upload...
Uploading files...
 3.61 MiB / 3.61 MiB  100.00% 2s
 ....
```

#### GitHub Actions
There is a workflow that will build, test, and deploy the app to cloud.gov when a pull request is created against the main branch or changes are pushed to master.

[1]: https://git-scm.com/
[2]: https://golang.org/
[3]: https://github.com/cloudfoundry/cli