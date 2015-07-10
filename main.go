package main

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
)

type Object struct {
	Id   string `json:"id"`
	Data interface{} `json:"data"`
	Version int `json:"version"`
}

type ObjectResource struct {
	objects map[string]Object
}

func (o ObjectResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/tables/usergraph/objects").
		Doc("Manage Objects").
		Consumes(restful.MIME_JSON, restful.MIME_XML).
		Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

	ws.Route(ws.GET("/{object-id}").To(o.findObject).
		// docs
		Doc("get an object").
		Operation("findObject").
		Param(ws.PathParameter("object-id", "identifier of the object").DataType("string")).
		Writes(Object{})) // on the response

	ws.Route(ws.GET("").To(o.listObjects).
		// docs
		Doc("get a list of objects").
		Operation("listObjects").
		Writes([]Object{})) // on the response

	ws.Route(ws.PUT("/{object-id}").To(o.updateObject).
		// docs
		Doc("update an object").
		Operation("updateObject").
		Param(ws.PathParameter("object-id", "identifier of the object").DataType("string")).
		ReturnsError(409, "Object was updated interim", Object{}).
		ReturnsError(404, "Object could not be found", nil).
		Reads(Object{})) // from the request

	ws.Route(ws.POST("").To(o.createObject).
		// docs
		Doc("create an object").
		Operation("createObject").
		Returns(201, "Object creted", Object{}).
		Reads(Object{})) // from the request
	container.Add(ws)
}

func main() {
	// to see what happens in the package, uncomment the following
	//restful.TraceLogger(log.New(os.Stdout, "[restful] ", log.LstdFlags|log.Lshortfile))
	o := ObjectResource{map[string]Object{}}
	wsContainer := restful.NewContainer()

	o.Register(wsContainer)

	// Optionally, you can install the Swagger Service which provides a nice Web UI on your REST API
	// You need to download the Swagger HTML5 assets and change the FilePath location in the config below.
	// Open http://localhost:8080/apidocs and enter http://localhost:8080/apidocs.json in the api input field.
	config := swagger.Config{
		WebServices:    wsContainer.RegisteredWebServices(), // you control what services are visible
		WebServicesUrl: "http://localhost:8090",
		ApiPath:        "/apidocs.json",

		// Optionally, specifiy where the UI is located
		SwaggerPath:     "/apidocs/",
		SwaggerFilePath: "/Users/kriaval/developer/swagger-ui"}
	swagger.RegisterSwaggerService(config, wsContainer)

	log.Printf("start listening on localhost:8090")
	server := &http.Server{Addr: ":8090", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
