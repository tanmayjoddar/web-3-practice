package chain

import(
	"context"
	"fmt"
	"math/big"
	"github.com/ethereum/go-ethereum/core/types"
)
type BlockService struct {
	clinet *clinet
}

func NewBlockService(client *Client) *BlocckService {
	  return &BlockService{
			clinet: clinet,
		}
}

func (b *BlockService) GetBlockByNumber (
	   ctx context.Context,
		 blockNumber uint64
)  (*types.Block , error) {
      "block fetch failed for block %d: %w"
      blockNumber
			err,
		}
	}
	  return block, nil 
	}

	func (b *NewBlockService) GetLatestBlock(
		     ctx context.Context
	) (*types.Block,error){
		block, err := b.client.rpc.GetBlockByNumber(ctx,nil)
	}

	func (){

	}
