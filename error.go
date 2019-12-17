package apileek

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "path"
    "runtime"
)

type bodyError struct {
    Error string `json:"error"`
}

type apiError struct {
    ApiError string
    StatusCode int
    ApiFunc string
}

type msgError struct {
    s string
}

func (e *msgError) Error() string {
    return e.s
}

func (e *apiError) Error() string {
    return fmt.Sprintf("Api: %s, http code: %d, api code: %s", e.ApiFunc, e.StatusCode, e.ApiError)
}

// Create a error with response information
func newApiError(
    resp *http.Response, // response of the request
    body string, // body of the request
) error {
    var obj bodyError

    if body == "[]" {
        obj = bodyError{Error: ""}
    } else {
        err := json.Unmarshal([]byte(body), &obj)
        if err != nil {
            log.Printf("Error to unmarshall body error: %s\nError:\n%v\n", body, err)
            return err
        }
    }
    pc, _, _, _ := runtime.Caller(1)
    f := runtime.FuncForPC(pc)

    return &apiError{
                ApiError: obj.Error,
                StatusCode: resp.StatusCode,
                ApiFunc: path.Base(f.Name()),
           }
}

// Create a simple message error
func newError(
    s string, // Error message
) error {
    return &msgError{
            s: s,
           }
}

