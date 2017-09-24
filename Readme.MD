# Bird Simple Web Server

Bird is a simple web server written in Golang.

You shouldn't use on your production server because of there are too many security risks.

# Usage

**Imports**

```go
import (
	"./anka"
	"./bird"
)
```

We use `anka` and `bird` to run server. You can set static file's directory with anka's **StaticDir(DirName)** method. So, you don't have to import anka package.

Bird is a main package to run server. These are important variables to get basic informations.

```go
// You can set base url like: 127.0.0.1 and port number like: "8081"
bird.Url // Get base URL
bird.Port // Get current port number, your project will run on this port.

bird.Type // Get request method
bird.Path // Get current url path. For example: 127.0.0.1/YOUR_PATH

bird.QueryList // You can get all query string parameters.

bird.Protocol // Get current HTTP procotol ~ 1.1 etc.

bird.StatusCode // Get response status code.

bird.RouteArray // Get all routes to get information.

```

## Methods:

These are important methods. For example:

**bird.Message** will add plain text string into webpage. Different from **bird.Template**

```go
bird.Message("Your Plain / Text Message")
```

If you will use Template method, you need template file. For example:

```go
p := Profile{Id: ID, User: Name, Path: Path, Type: Type, QueryList: bird.QueryList}

bird.Template("profile", p)
```

p is a optional struct. If you don't need it you can use `bird.Template("templateName", "")`.

**bird.Query** method will get single query string value. For example:

```go
// localhost/profile?id=1&name=ali

id := bird.Query("id")
name := bird.Query("name")
```

**bird.NotFound** important method to set Not Found messages. You can use plain text method or template. Example:

```go
// Message - Template name - Use Error Page with Template [yes / no]

bird.NotFound("Page not found!", "error/404.html", "yes")
```

**bird.Crow** this method most important method with Serve method. Because this method is responsible for setting routes and their methods.

Basic example with some methods:

**Basic Profile Struct**

```go
type Profile struct {
	Id string
	User string
	Path string
	Type string
	QueryList map[string][]string
}
```

**Codes**

```go
func HelloWorld() {
	bird.Message("Merhaba Dünya")
}

func ProfilePage() {

	Path := bird.Path
	Type := bird.Type

	ID := ""
	Name := ""

	if Type == "GET" && len(bird.QueryList) > 0 {
		ID = bird.Query("id")
		Name = bird.Query("name")
	}

	

	p := Profile{Id: ID, User: Name, Path: Path, Type: Type, QueryList: bird.QueryList}

	bird.Template("profile", p)
}

bird.Crow("/", HelloWorld)
bird.Crow("/profile", ProfilePage)
```

## Templates:

Bird uses Go's built-in template engine.

```go
Hello {{.User}}

Your ID: {{.Id}}

Current Path: {{.Path}}

Request Type: {{.Type}}

Query List: {{.QueryList}}
```

## All codes:

```go
package main

import (
	"./anka"
	"./bird"
)

type Profile struct {
	Id string
	User string
	Path string
	Type string
	QueryList map[string][]string
}

func HelloWorld() {
	bird.Message("Merhaba Dünya")
}

func ProfilePage() {

	Path := bird.Path
	Type := bird.Type

	ID := ""
	Name := ""

	if Type == "GET" && len(bird.QueryList) > 0 {
		ID = bird.Query("id")
		Name = bird.Query("name")
	}

	

	p := Profile{Id: ID, User: Name, Path: Path, Type: Type, QueryList: bird.QueryList}

	bird.Template("profile", p)
}

func main() {
	
	anka.StaticDir("static")

	bird.NotFound("Page not found!", "error/404.html", "yes")

	bird.Crow("/", HelloWorld)
	bird.Crow("/profile", ProfilePage)

	bird.Serve()

}
```

## Run

```bash
go run main.go
```

That's all :)