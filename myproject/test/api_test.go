package test

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestAPI(t *testing.T) {
    req, err := http.NewRequest("GET", "/user", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(GetUser)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Add more tests as needed
}
