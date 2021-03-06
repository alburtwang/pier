package client

import (
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/meshplus/pier/pkg/model"
)

// Client defines the interface that interacts with appchain
//go:generate mockgen -destination mock_client/mock_client.go -package mock_client -source client.go
type Client interface {
	// Start starts to listen appchain event
	Start() error

	// Stop stops client
	Stop() error

	// GetIBTP gets an interchain ibtp channel generated by client
	GetIBTP() chan *pb.IBTP

	// SubmitIBTP submits the interchain ibtp to appchain
	SubmitIBTP(*pb.IBTP) (*model.PluginResponse, error)

	// GetOutMessage gets interchain ibtp by index and target chain_id from broker contract
	GetOutMessage(to string, idx uint64) (*pb.IBTP, error)

	// GetInMessage gets receipt by index and source chain_id
	GetInMessage(from string, idx uint64) ([][]byte, error)

	// GetOutMeta gets an index map, which implicates the greatest index of
	// ingoing interchain txs for each source chain
	GetInMeta() (map[string]uint64, error)

	// GetOutMeta gets an index map, which implicates the greatest index of
	// outgoing interchain txs for each receiving chain
	GetOutMeta() (map[string]uint64, error)

	// GetReceiptMeta gets an index map, which implicates the greatest index of
	// executed callback txs for each receiving chain
	GetCallbackMeta() (map[string]uint64, error)

	// CommitCallback is a callback function when get receipt from bitxhub success
	CommitCallback(ibtp *pb.IBTP) error

	// Name gets name of blockchain from plugin
	Name() string

	// Type gets type of blockchain from plugin
	Type() string
}
