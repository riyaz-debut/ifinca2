// Package orders Related functions
package orders

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/chaincode/ifinca/pkg/core/status"
	"github.com/chaincode/ifinca/pkg/core/utils"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/s7techlab/cckit/router"
)

// CreateOrder create the Ifinca Coffee Order into the CouchDB
func CreateOrder(c router.Context) (interface{}, error) {
	// get the data from the request and parse it as structure
	data := c.Param(`data`).(Order)

	// set the default values for the fields
	data.DocType = utils.DocTypeOrder

	// Validate the inputed data
	err := data.Validate()
	if err != nil {
		if _, ok := err.(validation.InternalError); ok {
			return nil, err
		}
		return nil, status.ErrStatusUnprocessableEntity.WithValidationError(err.(validation.Errors))
	}

	// get the stub to use it for query and save
	stub := c.Stub()

	// check the order already exists or not
	queryString := fmt.Sprintf("{\"selector\":{\"order_no\":\"%s\",\"doc_type\":\"%s\"}}", data.OrderNo, utils.DocTypeOrder)
	alreadyExists, stateKey, _ := utils.Get(c, queryString, "")
	if alreadyExists != nil {
		// prepare the response body
		responseBody := utils.ResponseID{ID: stub.GetTxID()}
		// update the data and return the response
		return responseBody, c.State().Put(stateKey, data)
	}

	// prepare the response body
	responseBody := utils.ResponseID{ID: stub.GetTxID()}

	// Save the data and return the response
	return responseBody, c.State().Put(stub.GetTxID(), data)
}

// UpdateOrder update the Ifinca Coffee Order into the CouchDB
func UpdateOrder(c router.Context) (interface{}, error) {
	// get the data from the request and parse it as structure
	data := c.Param(`data`).(Order)

	// Validate the inputed data
	err := data.Validate()
	if err != nil {
		if _, ok := err.(validation.InternalError); ok {
			return nil, err
		}
		return nil, status.ErrStatusUnprocessableEntity.WithValidationError(err.(validation.Errors))
	}

	// set the default values for the fields
	data.DocType = utils.DocTypeOrder

	// get the stub to use it for query and save
	stub := c.Stub()

	// check the order exists or not
	queryString := fmt.Sprintf("{\"selector\":{\"order_no\":\"%s\",\"doc_type\":\"%s\"}}", data.OrderNo, utils.DocTypeOrder)
	orderDetails, stateKey, err := utils.Get(c, queryString, fmt.Sprintf("Order with order number: %s does not exists!", data.OrderNo))
	if orderDetails == nil {
		data.CreatedAt = time.Now()
		responseBody := utils.ResponseID{ID: stub.GetTxID()}
		// create order
		return responseBody, c.State().Put(stub.GetTxID(), data)
	}

	// ummarshal the byte array to structure
	orderdata := Order{}
	err = json.Unmarshal(orderDetails, &orderdata)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	data.CreatedAt = orderdata.CreatedAt
	responseBody := utils.ResponseID{ID: stub.GetTxID()}

	// return the response
	return responseBody, c.State().Put(stateKey, data)
}

// CreateSubOrders create the Ifinca Coffee Sub Orders into the CouchDB
func CreateSubOrders(c router.Context) (interface{}, error) {
	// get the data from the request and parse it as structure
	data := c.Param(`data`).(SubOrders)

	//check number of suborder
	if len(data.SubOrders) < 1 {
		return nil, status.ErrBadRequest.WithMessage(fmt.Sprintf("Invalid number of sub orders"))
	}

	//create a slice of subOrderIds to send in response
	subOrderIds := make([]string, 0, 0)
	// get the stub to use it for query and save
	stub := c.Stub()

	//insert asset one by one in blockchain
	for i, suborder := range data.SubOrders {

		// set the default values for the fields
		suborder.DocType = utils.DocTypeSubOrder
		suborder.CreatedAt = data.CreatedAt
		suborder.UpdatedAt = data.UpdatedAt

		// Validate the inputed data
		err := suborder.Validate()
		if err != nil {
			if _, ok := err.(validation.InternalError); ok {
				return nil, err
			}
			return nil, status.ErrStatusUnprocessableEntity.WithValidationError(err.(validation.Errors))
		}

		// check the order already exists or not
		orderQueryString := fmt.Sprintf("{\"selector\":{\"order_no\":\"%s\",\"doc_type\":\"%s\"}}", suborder.OrderNo, utils.DocTypeOrder)
		order_details, _, err1 := utils.Get(c, orderQueryString, fmt.Sprintf("Order with order no: %s does not exists!", suborder.OrderNo))
		if order_details == nil {
			return nil, err1
		}

		// check the sub order already exists or not
		subOrderQueryString := fmt.Sprintf("{\"selector\":{\"order_no\":\"%s\",\"supplier._id\":\"%s\",\"doc_type\":\"%s\"}}", suborder.OrderNo, suborder.Supplier.ID, utils.DocTypeSubOrder)
		alreadyExists, stateKey, _ := utils.Get(c, subOrderQueryString, "")
		if alreadyExists != nil {
			// update the data and return the response
			err := c.State().Put(stateKey, suborder)
			if err != nil {
				return nil, status.ErrInternal.WithError(err)
			}
			//append suborder id in slice after inserting suborder in blockchain
			subOrderIds = append(subOrderIds, stateKey)
		} else {
			// Save the data
			if suborderInsertErr := c.State().Put(stub.GetTxID()+strconv.Itoa(i), suborder); suborderInsertErr != nil {
				return nil, status.ErrInternal.WithError(suborderInsertErr)
			}
			//append suborder id in slice after inserting suborder in blockchain
			subOrderIds = append(subOrderIds, stub.GetTxID()+strconv.Itoa(i))
		}
	}

	responseBody := utils.ResponseID{ID: stub.GetTxID()}

	// return the response
	return responseBody, nil
}

// UpdateSubOrders update the Ifinca Coffee Sub-Order into the CouchDB
func UpdateSubOrders(c router.Context) (interface{}, error) {
	// get the data from the request and parse it as structure
	data := c.Param(`data`).(SubOrders)

	//check number of suborder
	if len(data.SubOrders) < 1 {
		return nil, status.ErrBadRequest.WithMessage(fmt.Sprintf("Invalid number of sub orders"))
	}
	// prepare the response body
	// responseBody := utils.ResponseMessage{Message: "success"}

	// get the stub to use it for query and save
	stub := c.Stub()

	//create a slice of subOrderIds to send in response
	subOrderIds := make([]string, 0, 0)

	//update sub order one by one in blockchain
	for i, suborder := range data.SubOrders {

		// Validate the inputed data
		err := suborder.Validate()
		if err != nil {
			if _, ok := err.(validation.InternalError); ok {
				return nil, err
			}
			return nil, status.ErrStatusUnprocessableEntity.WithValidationError(err.(validation.Errors))
		}

		suborder.DocType = utils.DocTypeSubOrder
		suborder.UpdatedAt = data.UpdatedAt

		// check the sub order exists or not
		queryString := fmt.Sprintf("{\"selector\":{\"order_no\":\"%s\",\"supplier._id\":\"%s\",\"doc_type\":\"%s\"}}", suborder.OrderNo, suborder.Supplier.ID, utils.DocTypeSubOrder)
		orderDetails, stateKey, _ := utils.Get(c, queryString, fmt.Sprintf("SubOrder with order no: %s does not exists!", suborder.OrderNo))
		if orderDetails == nil {
			suborder.CreatedAt = data.CreatedAt
			// create order
			err := c.State().Put(stub.GetTxID()+strconv.Itoa(i), suborder)
			if err != nil {
				return nil, status.ErrInternal.WithError(err)
			}
			subOrderIds = append(subOrderIds, stub.GetTxID()+strconv.Itoa(i))
		} else {
			// ummarshal the byte array to structure
			orderdata := SubOrder{}
			err = json.Unmarshal(orderDetails, &orderdata)
			if err != nil {
				return nil, status.ErrInternal.WithError(err)
			}

			// set the default values for the fields
			suborder.CreatedAt = orderdata.CreatedAt
			errUpdate := c.State().Put(stateKey, suborder)
			if errUpdate != nil {
				return nil, status.ErrInternal.WithError(errUpdate)
			}
			subOrderIds = append(subOrderIds, stateKey)
		}

	}
	// prepare the response body
	// responseBody := utils.ResponseIDs{IDs: subOrderIds}

	responseBody := utils.ResponseID{ID: stub.GetTxID()}

	// return the response
	return responseBody, nil
}

// GetHistoryForKey to get history
func GetHistoryForKey(c router.Context) (interface{}, error) {
	// get the data from the request and parse it as structure
	data := c.Param(`data`).(Key)

	// get the stub to use it for get history
	stub := c.Stub()
	var historyData bytes.Buffer
	historyData.WriteString("{\"history\": [")
	aArrayMemberAlreadyWritten := false

	historyIer, _ := stub.GetHistoryForKey(data.Key)
	for historyIer.HasNext() {
		// Add a comma before array members, suppress it for the first array member
		if aArrayMemberAlreadyWritten == true {
			historyData.WriteString(",")
		}

		modification, _ := historyIer.Next()
		historyData.WriteString(string(modification.Value))
		aArrayMemberAlreadyWritten = true
	}

	historyData.WriteString("]}")

	// return the response
	return historyData.Bytes(), nil
}

// GetOrders to get orders and suborders
func GetOrders(c router.Context) (interface{}, error) {
	// get the data from the request and parse it as structure
	data := c.Param(`data`).(OrderNo)

	// get the stub to use it for get orders
	stub := c.Stub()
	var orderData bytes.Buffer
	orderData.WriteString("{\"orders\": [")
	aArrayMemberAlreadyWritten := false
	orderQueryString := fmt.Sprintf("{\"selector\":{\"order_no\":\"%s\"}}", data.OrderNo)
	resultsIterator, err := stub.GetQueryResult(orderQueryString)
	// check if the query executed successfully?
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		// Add a comma before array members, suppress it for the first array member
		if aArrayMemberAlreadyWritten == true {
			orderData.WriteString(",")
		}

		order, _ := resultsIterator.Next()
		orderData.WriteString(string(order.Value))
		aArrayMemberAlreadyWritten = true
	}

	orderData.WriteString("]}")

	// return the response
	return orderData.Bytes(), nil
}
