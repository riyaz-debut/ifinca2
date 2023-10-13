// Package main Implements the Init & Invoke functions, Starts the chaincode
package main

import (
	"fmt"

	"github.com/chaincode/ifinca/pkg/orders"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/s7techlab/cckit/extensions/owner"
	"github.com/s7techlab/cckit/router"
	"github.com/s7techlab/cckit/router/param"
)

// Chaincode default chaincode implementation with router
type Chaincode struct {
	router *router.Group
}

// Init initializes chain code - sets chaincode "owner"
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	// delegate handling to router
	return cc.router.HandleInit(stub)
}

// Invoke - entry point for chain code invocations
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// delegate handling to router
	return cc.router.Handle(stub)
}

// New Define the Router
func New() *Chaincode {
	// create a new router instance
	r := router.New("Chaincode")
	chaincode := &Chaincode{r}

	// Handle the init/upgrade
	r.Init(invokeInit)

	// Other routes

	/***** Orders routes *****/

	r.Invoke(`createOrder`, orders.CreateOrder, param.Struct(`data`, &orders.Order{}))
	r.Invoke(`updateOrder`, orders.UpdateOrder, param.Struct(`data`, &orders.Order{}))
	r.Invoke(`createSubOrders`, orders.CreateSubOrders, param.Struct(`data`, &orders.SubOrders{}))
	r.Invoke(`updateSubOrders`, orders.UpdateSubOrders, param.Struct(`data`, &orders.SubOrders{}))
	r.Query(`getHistoryForKey`, orders.GetHistoryForKey, param.Struct(`data`, &orders.Key{}))
	r.Query(`getOrders`, orders.GetOrders, param.Struct(`data`, &orders.OrderNo{}))

	// return the routes
	return chaincode
}

// Invoked when the chaincode is instantiated or upgraded
func invokeInit(c router.Context) (interface{}, error) {
	return owner.SetFromCreator(c)
}

// Execution start point
func main() {
	if err := shim.Start(New()); err != nil {
		fmt.Printf("Error starting iFinca chaincode: %s", err)
	}
}
