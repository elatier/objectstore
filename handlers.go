package main

import (
	"github.com/emicklei/go-restful"
	"net/http"
	"strconv"
)

func (o ObjectResource) findObject(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("object-id")
	obj := o.objects[id]
	if len(obj.Id) == 0 {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "404: Object could not be found.")
		return
	}
	response.WriteEntity(obj)
}

func (o ObjectResource) listObjects(request *restful.Request, response *restful.Response) {
	values := make([]Object, len(o.objects))
	i := 0
	for _, value := range o.objects {
		values[i] = value
		i += 1
	}
	response.WriteEntity(values)
}

func (o *ObjectResource) createObject(request *restful.Request, response *restful.Response) {
	obj := new(Object)
	err := request.ReadEntity(obj)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	obj.Id = strconv.Itoa(len(o.objects) + 1) // simple id generation
	obj.Version = 0;
	o.objects[obj.Id] = *obj
	response.WriteHeader(http.StatusCreated)
	response.WriteEntity(obj)
}

func (o *ObjectResource) updateObject(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("object-id")
	curObj := o.objects[id]
	if len(curObj.Id) == 0 {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "404: Object could not be found:"+id)
		return
	}
	newObj := new(Object)
	err := request.ReadEntity(&newObj)
	newObj.Id = id
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	if newObj.Version != curObj.Version {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteHeader(http.StatusConflict)
		response.WriteEntity(curObj)
		return
	}
	newObj.Version += 1
	o.objects[newObj.Id] = *newObj
	response.WriteEntity(newObj)
}