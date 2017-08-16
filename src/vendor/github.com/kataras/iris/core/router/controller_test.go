// black-box testing
package router_test

import (
	"testing"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/core/router"

	"github.com/kataras/iris/httptest"
)

type testController struct {
	router.Controller
}

var writeMethod = func(c router.Controller) {
	c.Ctx.Writef(c.Ctx.Method())
}

func (c *testController) Get() {
	writeMethod(c.Controller)
}
func (c *testController) Post() {
	writeMethod(c.Controller)
}
func (c *testController) Put() {
	writeMethod(c.Controller)
}
func (c *testController) Delete() {
	writeMethod(c.Controller)
}
func (c *testController) Connect() {
	writeMethod(c.Controller)
}
func (c *testController) Head() {
	writeMethod(c.Controller)
}
func (c *testController) Patch() {
	writeMethod(c.Controller)
}
func (c *testController) Options() {
	writeMethod(c.Controller)
}
func (c *testController) Trace() {
	writeMethod(c.Controller)
}

type (
	testControllerAll struct{ router.Controller }
	testControllerAny struct{ router.Controller } // exactly same as All
)

func (c *testControllerAll) All() {
	writeMethod(c.Controller)
}

func (c *testControllerAny) All() {
	writeMethod(c.Controller)
}

func TestControllerMethodFuncs(t *testing.T) {
	app := iris.New()
	app.Controller("/", new(testController))
	app.Controller("/all", new(testControllerAll))
	app.Controller("/any", new(testControllerAny))

	e := httptest.New(t, app)
	for _, method := range router.AllMethods {

		e.Request(method, "/").Expect().Status(httptest.StatusOK).
			Body().Equal(method)

		e.Request(method, "/all").Expect().Status(httptest.StatusOK).
			Body().Equal(method)

		e.Request(method, "/any").Expect().Status(httptest.StatusOK).
			Body().Equal(method)
	}
}

type testControllerPersistence struct {
	router.Controller
	Data string `iris:"persistence"`
}

func (t *testControllerPersistence) Get() {
	t.Ctx.WriteString(t.Data)
}

func TestControllerPersistenceFields(t *testing.T) {
	data := "this remains the same for all requests"
	app := iris.New()
	app.Controller("/", &testControllerPersistence{Data: data})
	e := httptest.New(t, app)
	e.GET("/").Expect().Status(httptest.StatusOK).
		Body().Equal(data)
}

type testControllerBeginAndEndRequestFunc struct {
	router.Controller

	Username string
}

// called before of every method (Get() or Post()).
//
// useful when more than one methods using the
// same request values or context's function calls.
func (t *testControllerBeginAndEndRequestFunc) BeginRequest(ctx context.Context) {
	t.Username = ctx.Params().Get("username")
	// or t.Params.Get("username") because the
	// t.Ctx == ctx and is being initialized before this "BeginRequest"
}

// called after every method (Get() or Post()).
func (t *testControllerBeginAndEndRequestFunc) EndRequest(ctx context.Context) {
	ctx.Writef("done") // append "done" to the response
}

func (t *testControllerBeginAndEndRequestFunc) Get() {
	t.Ctx.Writef(t.Username)
}

func (t *testControllerBeginAndEndRequestFunc) Post() {
	t.Ctx.Writef(t.Username)
}

func TestControllerBeginAndEndRequestFunc(t *testing.T) {
	app := iris.New()
	app.Controller("/profile/{username}", new(testControllerBeginAndEndRequestFunc))

	e := httptest.New(t, app)
	usernames := []string{
		"kataras",
		"makis",
		"efi",
		"rg",
		"bill",
		"whoisyourdaddy",
	}
	doneResponse := "done"

	for _, username := range usernames {
		e.GET("/profile/" + username).Expect().Status(httptest.StatusOK).
			Body().Equal(username + doneResponse)
		e.POST("/profile/" + username).Expect().Status(httptest.StatusOK).
			Body().Equal(username + doneResponse)
	}

}
