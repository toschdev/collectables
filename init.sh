#!/bin/bash
rm -rf ~/.colld
rm -rf ~/.collcli

colld init collectables --chain-id collectables-chain

collcli config keyring-backend test

collcli keys add collector1
collcli config indent true
collcli config output json
collcli config trust-node true
collcli config chain-id collectables-chain
colld add-genesis-account $(collcli keys show collector1 -a) 100000000stake
colld gentx --name collector1 --keyring-backend test
colld collect-gentxs                                                 
colld validate-genesis                                               
colld start        