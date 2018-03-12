# web_server_template

### 1. Set your $GOPATH in ~/.zshrc
(1) Go to web_server_template
(2) Type $pwd
(3) Copy the path to ~/.zshrc like this:
  export GOPATH=$XXX/XXX/XXX/XXX/web_server_template

### 2. Explanation of each files
#### 2.1. server.go
If you have a server, type "$ git clone git@github.com:TakahitoMotoki/web_server_template.git" in your command line.
Then, type "$ go run server.go" to start web_server_template.

You have to change Routing Zone if you need.
// MARK: - Routing Zone Start
// MARK: - Routing Zone End

#### 2.2. src/handler.go
This file helps you to make responses according to the URL from clients.

You have to change all method if you need.
