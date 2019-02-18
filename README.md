## Welcome to the front-end developer code test!

### This is a small golang binary that will serve up some static data for the code test.

#### to run the go binary
* install golang (https://golang.org/doc/install)
* clone the repo into a sensible place for your GOPATH
* `go run main.go`

< Note: if you have trouble getting the golang app up and running, please contact us right away so we can unblock you >

#### sample api calls

```
curl http://127.0.0.1:2133/nodes -d '{}'
curl http://127.0.0.1:2133/nodes -d '{"filters": [ { "key": "last_scan_status", "value": "failed" } ] }'
curl http://127.0.0.1:2133/nodes -d '{"filters": [ { "key": "name", "value": "my test node" } ] }'
```

## The code test project: what to do

Build a GUI to display and filter the information returned by the `nodes` endpoint.
This should include:
* a table
* filter search bar

What we're looking at when evaluating your code project:
* data flow models
* accessibility
* MORE??

You can choose whatever framework(s) you'd like, including none.
Please include some details on how to run your code. This can be in a README, Makefile, whatever you'd like.
As part of the interview process, you will be reviewing your code with two engineers from Chef.

