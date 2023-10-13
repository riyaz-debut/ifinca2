// Package utils Common Strcutres, Constants etc that are being used in other packages
package utils

import (
	"time"

	"github.com/chaincode/ifinca/pkg/core/status"

	"github.com/s7techlab/cckit/router"
)

// MetaData Strcuture: Contains the common fields which are used in all other Structures
type MetaData struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	DocType   string    `json:"doc_type"`
}

// ResponseID is used to return the response which contains only one ID field
type ResponseID struct {
	ID string `json:"id"`
}

// ResponseIDs is used to return the response which contains only one ID field
type ResponseIDs struct {
	IDs []string `json:"ids"`
}

// ResponseMessage is used to return the response which contains only one message field
type ResponseMessage struct {
	Message string `json:"message"`
}

// Constants DocTypes the document which are stored inside the couchdb
const (
	DocTypeOrder    string = "order"    // For orders
	DocTypeSubOrder string = "suborder" // For subOrders
)

// Get Finds the record by ID
func Get(c router.Context, query string, message string) ([]byte, string, error) {
	stub := c.Stub()
	// excecute the query
	resultsIterator, err := stub.GetQueryResult(query)

	// check if the query executed successfully?
	if err != nil {
		return nil, "", status.ErrInternal.WithError(err)
	}
	defer resultsIterator.Close()

	// query has returned the results?
	if !resultsIterator.HasNext() {
		return nil, "", status.ErrNotFound.WithMessage(message)
	}

	// fetch the data and marshal it into struct
	queryResponse, err := resultsIterator.Next()
	if err != nil {
		return nil, "", status.ErrInternal.WithError(err)
	}

	return queryResponse.Value, queryResponse.Key, nil
}
