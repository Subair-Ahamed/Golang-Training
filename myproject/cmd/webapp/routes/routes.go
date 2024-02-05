package routes

import (
    "net/http"
    "myproject/api/controllers"
)

func NewRouter() *http.ServeMux {
    mux := http.NewServeMux()

    mux.HandleFunc("/user", controllers.GetUser)
    mux.HandleFunc("/user/update", controllers.UpdateUser)

    return mux
}
