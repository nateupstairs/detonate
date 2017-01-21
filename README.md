# Detonate

### Go simple JSON HTTP errors

This exists as I love the way the node library hapijs/boom handles errors and I miss it while working in Go. This project will never be inteded to be as complete as boom. This will mostly be a way to abstract away the process of creating simple JSON error responses and writing it to a given writer with correct status code.

an example 500 error:

```go
func route(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	d := detonate.Internal("Oops")
    d.trigger(w)
    return
}
```

generates

```json
{
	"code":500,
    "error":"Internal Server Error",
    "message":"Oops"
}
```

an example 404 error:

```go
func route(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	d := detonate.NotFound("Image Missing")
    d.trigger(w)
    return
}
```

generates

```json
{
	"code":404,
    "error":"Not Found",
    "message":"Image Missing"
}
```

You can also generate an error manually:

```go
func route(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	d := detonate.Create(500, "Weird Error", "Something Is Off")
    ...
    d.trigger(w)
    return
}
```
```json
{
	"code":500,
    "error":"Weird Error",
    "message":"Something Is Off"
}
```

you can add validations if you want to give specific feedback for user data:

```go
func route(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	d := detonate.BadData("Information Invalid")
    v := detonate.CreateValidation("testkey", "Test", "maybe don't do it wrong?")
    d.AddValidation(v)
    d.trigger(w)
    return
}
```

generates

```json
{
    "code":422,
    "error":"Unprocessable Entity",
    "message":"Information Invalid",
    "validations":[
        {
            "key":"testkey",
            "error":"Test",
            "message":"maybe don't do it wrong?"
        }
    ]
}
```