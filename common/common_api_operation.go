package common

import (
	"fmt"
)

func API_URL(operation string, namespace string, name string) (conten string) {
	//list or watch objects of kind Endpoints
	howop := operation
	componets_one := "componentstatuses"
	componets_two := "componentstatuses/" + name

	endpoints_one := "namespaces/" + namespace + "/endpoints"
	endpoints_two := "namespaces/" + namespace + "/endpoints/" + name
	endpoints_three := "endpoints"

	service_one := "namespaces/" + namespace + "/services"
	service_two := "namespaces/" + namespace + "/services/" + name
	service_three := "services"

	node_one := "nodes"
	node_two := "nodes/" + name

	replication_one := "replicationcontrollers"
	replication_two := "namespaces/" + namespace + "/replicationcontrollers/" + name

	pod_one := "pods"
	pod_two := "namespaces/" + namespace + "/pods/" + name

	switch howop {
	//list objects of kind ComponentStatus
	//GET /api/v1/componentstatuses
	case "ComponetStatusList":
		conten = componets_one
	//read the specified ComponentStatus
	//GET /api/v1/componentstatuses/{name}
	case "SpeccomponetsStatus":
		conten = componets_two

		//list or watch objects of kind Endpoints
		//GET /api/v1/endpoints
	case "ListEndpoints":
		conten = endpoints_three

	//list or watch objects of kind Endpoints
	case "listOrWatch":
		conten = endpoints_one

		//delete collection of Endpoints
	case "deleteCollectionOfEndpoints":
		conten = endpoints_one

		// create a Endpoints
	case "createAnEndpoints":
		conten = endpoints_one

		//read the specified Endpoints
	case "readSpecifiedEndpoints":
		conten = endpoints_two

	//delete a Endpoints
	case "deleteAnEndpoints":
		conten = endpoints_two

		//	replace the specified Endpoints
	case "replaceSpecifiedEndpoints":
		conten = endpoints_two

		//	partially update the specified Endpoints
	case "paritallUpdateSpecifiedEndpoints":
		conten = endpoints_two

		//list all namespace service
		//Get /api/v1/service
	case "NamespaceServiceList":
		conten = service_three

		//list or watch objects of kind Service
		//GET /api/v1/namespaces/{namespace}/services
	case "listOneNamespaceService":
		conten = service_one

		//create a Service
		// POST /api/v1/namespaces/{namespace}/services
	case "createAService":
		conten = service_one

		//read the specified Service
	// GET /api/v1/namespaces/{namespace}/services/{name}
	case "readSpecService":
		conten = service_two

		// replace the specified Service
	// PUT /api/v1/namespaces/{namespace}/services/{name}
	case "replaceSpecService":
		conten = service_two

	//delete a Service
	// DELETE /api/v1/namespaces/{namespace}/services/{name}
	case "deleteAService":
		conten = service_two

	//partially update the specified Service
	// PATCH /api/v1/namespaces/{namespace}/services/{name}
	case "updateSpecService":
		conten = service_two

		// connect GET requests to proxy of Service
		// GET /api/v1/namespaces/{namespace}/services/{name}/proxy

		// connect PUT requests to proxy of Service
		// PUT /api/v1/namespaces/{namespace}/services/{name}/proxy

		// connect DELETE requests to proxy of Service
		// DELETE /api/v1/namespaces/{namespace}/services/{name}/proxy

		// connect POST requests to proxy of Service
		// POST /api/v1/namespaces/{namespace}/services/{name}/proxy

		// connect GET requests to proxy of Service
		// GET /api/v1/namespaces/{namespace}/services/{name}/proxy/{path}

		// connect PUT requests to proxy of Service
		// PUT /api/v1/namespaces/{namespace}/services/{name}/proxy/{path}

		// connect DELETE requests to proxy of Service
		// DELETE /api/v1/namespaces/{namespace}/services/{name}/proxy/{path}

		// connect POST requests to proxy of Service
		// POST /api/v1/namespaces/{namespace}/services/{name}/proxy/{path}

		// read status of the specified Service
		// GET /api/v1/namespaces/{namespace}/services/{name}/status

		// replace status of the specified Service
		// PUT /api/v1/namespaces/{namespace}/services/{name}/status

		// partially update status of the specified Service
		// PATCH /api/v1/namespaces/{namespace}/services/{name}/status

		//list all replicationcontroller
	case "replication_listall":
		conten = replication_one
		// list or watch objects of kind ReplicationController
		// GET /api/v1/namespaces/{namespace}/replicationcontrollers

		// delete collection of ReplicationController
		// DELETE /api/v1/namespaces/{namespace}/replicationcontrollers

		// create a ReplicationController
		// POST /api/v1/namespaces/{namespace}/replicationcontrollers

		// read the specified ReplicationController
		// GET /api/v1/namespaces/{namespace}/replicationcontrollers/{name}
	case "ReadSpecReplication":
		conten = replication_two
		// replace the specified ReplicationController
		// PUT /api/v1/namespaces/{namespace}/replicationcontrollers/{name}

		// delete a ReplicationController
		// DELETE /api/v1/namespaces/{namespace}/replicationcontrollers/{name}

		// partially update the specified ReplicationController
		// PATCH /api/v1/namespaces/{namespace}/replicationcontrollers/{name}

		// read scale of the specified Scale
		// GET /api/v1/namespaces/{namespace}/replicationcontrollers/{name}/scale

		// replace scale of the specified Scale
		// PUT /api/v1/namespaces/{namespace}/replicationcontrollers/{name}/scale

		// partially update scale of the specified Scale
		// PATCH /api/v1/namespaces/{namespace}/replicationcontrollers/{name}/scale

		// read status of the specified ReplicationController
		// GET /api/v1/namespaces/{namespace}/replicationcontrollers/{name}/status

		// replace status of the specified ReplicationController
		// PUT /api/v1/namespaces/{namespace}/replicationcontrollers/{name}/status

		// partially update status of the specified ReplicationController
		// PATCH /api/v1/namespaces/{namespace}/replicationcontrollers/{name}/status
		// read the specified Namespace
		// GET /api/v1/namespaces/{name}

		// replace the specified Namespace
		// PUT /api/v1/namespaces/{name}

		// delete a Namespace
		// DELETE /api/v1/namespaces/{name}

		// partially update the specified Namespace
		// PATCH /api/v1/namespaces/{name}

		// replace finalize of the specified Namespace
		// PUT /api/v1/namespaces/{name}/finalize

		// read status of the specified Namespace
		// GET /api/v1/namespaces/{name}/status

		// replace status of the specified Namespace
		// PUT /api/v1/namespaces/{name}/status

		// partially update status of the specified Namespace
		// PATCH /api/v1/namespaces/{name}/status

		// list or watch objects of kind Node
		// GET /api/v1/nodes
	case "Nodelist":
		conten = node_one
		// delete collection of Node
		// DELETE /api/v1/nodes

		// create a Node
		// POST /api/v1/nodes

		// read the specified Node
		// GET /api/v1/nodes/{name}
	case "SpecNode":
		conten = node_two
		// replace the specified Node
		// PUT /api/v1/nodes/{name}

		// delete a Node
		// DELETE /api/v1/nodes/{name}

		// partially update the specified Node
		// PATCH /api/v1/nodes/{name}

		// connect GET requests to proxy of Node
		// GET /api/v1/nodes/{name}/proxy

		// connect PUT requests to proxy of Node
		// PUT /api/v1/nodes/{name}/proxy

		// connect DELETE requests to proxy of Node
		// DELETE /api/v1/nodes/{name}/proxy

		// connect POST requests to proxy of Node
		// POST /api/v1/nodes/{name}/proxy

		// connect GET requests to proxy of Node
		// GET /api/v1/nodes/{name}/proxy/{path}

		// connect PUT requests to proxy of Node
		// PUT /api/v1/nodes/{name}/proxy/{path}

		// connect DELETE requests to proxy of Node
		// DELETE /api/v1/nodes/{name}/proxy/{path}

		// connect POST requests to proxy of Node
		// POST /api/v1/nodes/{name}/proxy/{path}

		// read status of the specified Node
		// GET /api/v1/nodes/{name}/status

		// replace status of the specified Node
		// PUT /api/v1/nodes/{name}/status

		// partially update status of the specified Node
		// PATCH /api/v1/nodes/{name}/status

		// list or watch objects of kind PersistentVolumeClaim
		// GET /api/v1/persistentvolumeclaims

		// list or watch objects of kind PersistentVolume
		// GET /api/v1/persistentvolumes

		// delete collection of PersistentVolume
		// DELETE /api/v1/persistentvolumes

		// create a PersistentVolume
		// POST /api/v1/persistentvolumes

		// read the specified PersistentVolume
		// GET /api/v1/persistentvolumes/{name}

		// replace the specified PersistentVolume
		// PUT /api/v1/persistentvolumes/{name}

		// delete a PersistentVolume
		// DELETE /api/v1/persistentvolumes/{name}

		// partially update the specified PersistentVolume
		// PATCH /api/v1/persistentvolumes/{name}

		// read status of the specified PersistentVolume
		// GET /api/v1/persistentvolumes/{name}/status

		// replace status of the specified PersistentVolume
		// PUT /api/v1/persistentvolumes/{name}/status

		// partially update status of the specified PersistentVolume
		// PATCH /api/v1/persistentvolumes/{name}/status

		// list or watch objects of kind Pod
		// GET /api/v1/pods
	case "Pod_list":
		conten = pod_one

		//GET /api/v1/namespaces/{namespace}/pods/{name}
	case "Spec_pod":
		conten = pod_two

		// list or watch objects of kind PodTemplate
		// GET /api/v1/podtemplates

		// proxy GET requests to Pod
		// GET /api/v1/proxy/namespaces/{namespace}/pods/{name}

		// proxy PUT requests to Pod
		// PUT /api/v1/proxy/namespaces/{namespace}/pods/{name}

		// proxy DELETE requests to Pod
		// DELETE /api/v1/proxy/namespaces/{namespace}/pods/{name}

		// proxy POST requests to Pod
		// POST /api/v1/proxy/namespaces/{namespace}/pods/{name}

		// proxy GET requests to Pod
		// GET /api/v1/proxy/namespaces/{namespace}/pods/{name}/{path}

		// proxy PUT requests to Pod
		// PUT /api/v1/proxy/namespaces/{namespace}/pods/{name}/{path}

		// proxy DELETE requests to Pod
		// DELETE /api/v1/proxy/namespaces/{namespace}/pods/{name}/{path}

		// proxy POST requests to Pod
		// POST /api/v1/proxy/namespaces/{namespace}/pods/{name}/{path}

		// proxy GET requests to Service
		// GET /api/v1/proxy/namespaces/{namespace}/services/{name}

		// proxy PUT requests to Service
		// PUT /api/v1/proxy/namespaces/{namespace}/services/{name}

		// proxy DELETE requests to Service
		// DELETE /api/v1/proxy/namespaces/{namespace}/services/{name}

		// proxy POST requests to Service
		// POST /api/v1/proxy/namespaces/{namespace}/services/{name}

		// proxy GET requests to Service
		// GET /api/v1/proxy/namespaces/{namespace}/services/{name}/{path}

		// proxy PUT requests to Service
		// PUT /api/v1/proxy/namespaces/{namespace}/services/{name}/{path}

		// proxy DELETE requests to Service
		// DELETE /api/v1/proxy/namespaces/{namespace}/services/{name}

		// proxy POST requests to Service
		// POST /api/v1/proxy/namespaces/{namespace}/services/{name}

		// proxy GET requests to Service
		// GET /api/v1/proxy/namespaces/{namespace}/services/{name}/{path}

		// proxy PUT requests to Service
		// PUT /api/v1/proxy/namespaces/{namespace}/services/{name}/{path}

	}
	fmt.Println("common_api.conten", conten)
	return conten
}
