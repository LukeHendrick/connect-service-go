# Connect Service GO!

A Connect API Mock Service built with Golang

## Download

A Docker Image is available as:
```bash
lukehendrick/connect-service-go
```

## Usage

Start the container with:
```bash
docker run $HOST_PORT:9000 lukehendrick/connect-service-go
```

Where `$HOST_PORT` is the local port you want to expose

## Contributing

This is very much a WIP, eventually we want to recreate all routes listed **[here](https://docs.aws.amazon.com/connect/latest/APIReference/API_Operations.html)**

When creating a new route, the base path should match the folder name. For example, the `ListContactFlows` method request is:
```bash
GET /contact-flows-summary/InstanceId?contactFlowTypes=ContactFlowTypes&maxResults=MaxResults&nextToken=NextToken HTTP/1.1
```
So this method is defined inside of the contactflowssummary folder (to match go pkg naming standards)

Routes are defined in a similarly named package inside the path folder:
```golang
package contactflowssummary

import (
	"github.com/go-chi/chi"
)

// Routes for the ContactFlow Operations
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{instanceId}", ListContactFlows)
	return router
}
```
Where each HTTP Method/Path will be added to the list. These routes are then mounted inside the `routes` method in `main.go`
```golang
router.Route("/", func(r chi.Router) {
  r.Mount("/contact-flows-summary", contactflowssummary.Routes())
  //...additional routes here
})
```

Tests should reside in the same folder and follow the `FILE_test.go` naming convention

Please raise any issues you find on the GitHub Repository [here](https://github.com/LukeHendrick/connect-service-go.git)