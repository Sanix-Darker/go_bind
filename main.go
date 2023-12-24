package main

/*
#cgo pkg-config: python3
#define Py_LIMITED_API
#include <Python.h>
*/
import "C"
import (
	"context"
	"fmt"
	"log"
	"unsafe"

	v0 "github.com/authzed/authzed-go/proto/authzed/api/v0"
	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
)

//export InitAuthZClient
func InitAuthZClient(host *C.char, bearerToken *C.char) unsafe.Pointer {
	systemCerts, err := grpcutil.WithSystemCerts(
		grpcutil.VerifyCA,
	)

	if err != nil {
		log.Fatalf("unable to load system CA certificates: %s", err)
	}

	client, err := authzed.NewClient(
		C.GoString(host),
		systemCerts,
		grpcutil.WithBearerToken(C.GoString(bearerToken)),
	)
	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}
	log.Println(client)

	return unsafe.Pointer(client)
}

//export CheckPermission
func CheckPermission(clientPtr unsafe.Pointer, namespace *C.char, objectID *C.char) (bool, error) {
	client := (*authzed.Client)(clientPtr)

	permission := &v0.RelationTuple{
		ObjectAndRelation: &v0.ObjectAndRelation{
			Namespace: C.GoString(namespace),
			ObjectId:  C.GoString(objectID),
			Relation:  "can_read",
		},
	}

	checkRequest := &v1.CheckPermissionRequest{
		Permission: permission,
	}

	response, err := client.CheckPermission(context.Background(), checkRequest)
	if err != nil {
		return false, err
	}

	fmt.Printf("response :: %v", response)

	return true, nil
}

//export UpdateRelation
func UpdateRelation(clientPtr unsafe.Pointer, namespace *C.char, objectID *C.char) error {
	client := (*authzed.Client)(clientPtr)

	newRelation := &v0.RelationTuple{
		ObjectAndRelation: &v0.ObjectAndRelation{
			Namespace: C.GoString(namespace),
			ObjectId:  C.GoString(objectID),
			Relation:  "can_write",
		},
	}

	updateRequest := &v1.WriteRelationTupleRequest{
		Updates: []*v1.RelationTupleUpdate{{
			Tuple:  newRelation,
			Action: v1.RelationTupleUpdate_MODIFY,
		}},
	}

	_, err := client.WriteRelationTuple(context.Background(), updateRequest)
	return err
}

func main() {}
