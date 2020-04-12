package service

import (
	"fmt"

	"gitlab.com/vocdoni/go-dvote/census"
	"gitlab.com/vocdoni/go-dvote/chain"
	"gitlab.com/vocdoni/go-dvote/chain/ethevents"
	"gitlab.com/vocdoni/go-dvote/log"
)

// EthEvents service registers on the Ethereum smart contract specified in ethProcDomain, the provided event handlers
// we3host and w3port must point to a working web3 websocket endpoint
// If subscribe is enabled the service will also subscribe for new blocks
func EthEvents(ethProcDomain, ensRegistryAddr, w3host string, w3port int, startBlock int64, endBlock int64, subscribe bool,
	cm *census.Manager, evh []ethevents.EventHandler) error {
	// TO-DO remove cm (add it on the eventHandler instead)
	log.Infof("creating ethereum events service")

	contractAddr, err := chain.VotingProcessAddress(
		ensRegistryAddr, ethProcDomain, fmt.Sprintf("ws://%s:%d", w3host, w3port))
	if err != nil || contractAddr == "" {
		return fmt.Errorf("cannot get voting process contract: %s", err)
	} else {
		log.Infof("loaded voting contract at address: %s", contractAddr)
	}

	ev, err := ethevents.NewEthEvents(contractAddr, nil, fmt.Sprintf("ws://%s:%d", w3host, w3port), cm)
	if err != nil {
		return fmt.Errorf("couldn't create ethereum events listener: (%s)", err)
	}
	for _, e := range evh {
		ev.AddEventHandler(e)
	}
	go func() {
		go ev.ReadEthereumEventLogs(startBlock, endBlock)
		log.Infof("subscribing to new ethereum events from block %d", endBlock)
		ev.SubscribeEthereumEventLogs()
	}()

	return nil
}