#!/bin/bash

#Create and join channel
export CHANNEL_NAME=mychannel
export CHANECODE_VERSION=1.8

# the channel.tx file is mounted in the channel-artifacts directory within your CLI container
# as a result, we pass the full path for the file
# we also pass the path for the orderer ca-cert in order to verify the TLS handshake
# be sure to export or replace the $CHANNEL_NAME variable appropriately
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/ifinca.co/users/Admin@ifinca.co/msp
export CORE_PEER_ADDRESS=peer0.ifinca.co:7051
export CORE_PEER_LOCALMSPID="ifincaMSP"
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/ifinca.co/peers/peer0.ifinca.co/tls/ca.crt

#Install chaincode
peer chaincode install -n my_chaincode -v $CHANECODE_VERSION -p github.com/chaincode/ifinca/cmd

sleep 4

peer chaincode upgrade -o orderer0.ifinca.co:7050 -C mychannel -n my_chaincode -v $CHANECODE_VERSION -c '{"Args":[]}' -P "OR ('ifincaMSP.member')"
