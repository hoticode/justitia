package node

import (
	"github.com/DSiSc/craft/types"
	"github.com/DSiSc/justitia/common"
	"github.com/DSiSc/justitia/tools/events"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

// mock a transaction
func mock_transactions(num int) []*types.Transaction {
	to := make([]types.Address, num)
	for m := 0; m <= num; m++ {
		for j := 0; j < types.AddressLength; j++ {
			to[m][j] = byte(m)
		}
	}
	amount := new(big.Int)
	txList := make([]*types.Transaction, 0)
	for i := 1; i <= num; i++ {
		tx := common.NewTransaction(uint64(i), to[i], amount.SetUint64(uint64(i)), uint64(i), amount, nil, to[0])
		txList = append(txList, tx)
	}
	return txList
}

func TestNewNode(t *testing.T) {
	assert := assert.New(t)

	service, err := NewNode()
	assert.Nil(err)
	assert.NotNil(service)

	nodeService := service.(*Node)
	assert.NotNil(nodeService.txpool)
	assert.NotNil(nodeService.participates)
	assert.NotNil(nodeService.txSwitch)
	assert.NotNil(nodeService.blockSwitch)
	assert.NotNil(nodeService.consensus)
	assert.NotNil(nodeService.config)
	assert.NotNil(nodeService.role)
	assert.Nil(nodeService.producer)
	assert.Nil(nodeService.validator)
	event := types.GlobalEventCenter.(*events.Event)
	assert.Equal(3, len(event.Subscribers))
}

func TestNode_Start(t *testing.T) {
	assert := assert.New(t)
	service, err := NewNode()
	assert.Nil(err)
	assert.NotNil(service)
	go func() {
		service.Start()
		nodeService := service.(*Node)
		assert.NotNil(nodeService.rpcListeners)
		assert.Equal(1, len(nodeService.rpcListeners))
		service.Wait()
	}()
	service.Stop()
}

func TestEventRegister(t *testing.T) {
	EventRegister()
	event := types.GlobalEventCenter.(*events.Event)
	assert.Equal(t, 3, len(event.Subscribers))
}
