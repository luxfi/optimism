// Package optimism provides a rollup plugin for Optimism (OP Stack) integration
package optimism

import (
   "context"
   "github.com/ethereum-optimism/optimism/op-node"
   "github.com/luxfi/geth/mode"
   "github.com/luxfi/geth/rollup"
   "github.com/luxfi/geth/utils/logging"
   "github.com/luxfi/geth/rpc"
)

// Plugin implements the RollupPlugin interface for Optimism.
type Plugin struct {
   config *mode.RollupConfig
   node   *op_node.RollupNode
}

// Name returns the rollup type identifier.
func (p *Plugin) Name() string {
   return "optimism"
}

// RegisterGenesis initializes the Optimism chain genesis.
func (p *Plugin) RegisterGenesis(genesisPath string) error {
   return op_node.InitGenesis(genesisPath)
}

// StartSequencer starts the Optimism sequencer node.
func (p *Plugin) StartSequencer(ctx context.Context) error {
   if p.config.SequencerEndpoint == "" {
       return nil
   }
   // Create and start the rollup node
   node, err := op_node.NewRollupNode(op_node.RollupNodeConfig{
       L1NodeURL:     p.config.L1Endpoint,
       SequencerURL:  p.config.SequencerEndpoint,
   })
   if err != nil {
       return err
   }
   p.node = node
   return p.node.Start(ctx)
}

// RegisterRPC registers Optimism-specific RPC namespaces.
func (p *Plugin) RegisterRPC(server *rpc.Server) {
   op_node.RegisterRPC(server)
}

// NewPlugin constructs a new Optimism rollup plugin using given config.
func NewPlugin(config *mode.RollupConfig) rollup.Plugin {
   return &Plugin{config: config}
}