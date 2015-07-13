# objectstore
Build using "go build" and run using "./objectstore".
Swagger spec file is available at http://localhost:8090/apidocs.json by default.
Swagger UI needs to be downloaded separately from https://github.com/swagger-api/swagger-ui to view it.

Default API base URL: http://localhost:8090/tables/usergraph/objects

This service provides generic object storage in format:

type Object struct {
	Id      string      `json:"id"`
	Data    interface{} `json:"data"`
	Version int         `json:"version"`
}

Data - can contain generic json compatible data. It is currently used by https://github.com/elatier/usergraph for persistence.