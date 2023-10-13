#!/bin/bash

# #Generate certificates
# cryptogen generate --config=crypto-config.yaml

# export FABRIC_CFG_PATH=$PWD

# #Create Channel Artifacts
# #Genesis Block
# configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block

# #Channel.tx
# export CHANNEL_NAME=mychannel && configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME

# #Org1AnchorPeer.tx
# export CHANNEL_NAME=mychannel && configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/debutMSPanchors.tx -asOrg debutMSP -channelID $CHANNEL_NAME

# #Org2AnchorPeer.tx
# export CHANNEL_NAME=mychannel && configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/googleMSPanchors.tx -asOrg googleMSP -channelID $CHANNEL_NAME

#Network Up
docker-compose -f ./docker-compose-cli.yaml up -d

# docker exec -it cli bash


