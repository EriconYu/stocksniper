package main

import (
	"net"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx context.Context) {
		ctx.Writef("Hello from the server")
	})

	app.Get("/mypath", func(ctx context.Context) {
		ctx.Writef("Hello from %s", ctx.Path())
	})

	// create any custom tcp listener, unix sock file or tls tcp listener.
	l, err := net.Listen("tcp4", ":8080")
	if err != nil {
		panic(err)
	}

	// use of the custom listener
	app.Run(iris.Listener(l))
}
