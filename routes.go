package main

import (
  "net/http"

  "github.com/gorilla/mux"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return router
}

var routes = Routes{
    Route{
        "Swagger",
        "GET",
        "/v1/swagger.json",
        Swagger,
    },
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "LocationsIndex",
        "GET",
        "/v1/locations",
        LocationsIndex,
    },
    Route{
        "TrendsIndex",
        "GET",
        "/v1/trends/{location}/{term}",
        TrendsIndex,
    },
    Route{
        "TrendsRouteIndex",
        "GET",
        "/v1/trends/{location}",
        TrendsRouteIndex,
    },
    Route{
        "WebTrendsIndex",
        "GET",
        "/web/trends/{location}",
        WebTrendsRouteIndex,
    },
    Route{
        "WebTrendsIndex",
        "GET",
        "/web/trends/{location}/{term}",
        WebTrendsIndex,
    },
    Route{
        "AdminIndex",
        "GET",
        "/admin/",
        AdminIndex,
    },
    Route{
        "AdminBuildDatabase",
        "GET",
        "/admin/builddatabase",
        AdminBuildDatabase,
    },
    Route{
        "AdminBuildData",
        "GET",
        "/admin/builddata",
        AdminBuildData,
    },
    Route{
        "AdminBuildSeeds",
        "GET",
        "/admin/buildseeds",
        AdminBuildSeeds,
    },
    Route{
        "AdminSeeds",
        "GET",
        "/admin/seeds",
        AdminSeeds,
    },
}