package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx context.Context) {
		ctx.HTML("<h1>Hello World!</h1>")
	})

	err := app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		// do something
	}
	// same as:
	// err := app.Run(iris.Addr(":8080"))
	// if err != nil && (err != iris.ErrServerClosed || err.Error() != iris.ErrServerClosed.Error()) {
	//     [...]
	// }
}
