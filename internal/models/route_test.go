package models

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestNew(t *testing.T) {
	// sample new
	resp, err := New(1, "test route name", 5.5)
	expected := Route{
		ID:        resp.ID,
		UserID:    1,
		RouteName: "test route name",
		Length:    5.5,
	}
	assert.Assert(t, err == nil)
	assert.Assert(t, resp == expected)

	// negative lenght
	resp, err = New(1, "test route name", -5.5)
	assert.Assert(t, err.Error() == "route lenght must be greater than zero")
	assert.Assert(t, resp == Route{})
}

func TestRouteTime(t *testing.T) {
	// sample routeTime
	rt := Route{1, 1, "test route name", 5}
	resp, err := rt.RouteTime(rt, 2)
	assert.Assert(t, err == nil)
	assert.Assert(t, resp == 2.5)

	// negative speed
	rt = Route{1, 1, "test route name", 5}
	resp, err = rt.RouteTime(rt, -2)
	fmt.Println("resp", resp)
	fmt.Println("err", err)
	assert.Assert(t, err.Error() == "speed must be greater than zero")
	assert.Assert(t, resp == 0)
}
