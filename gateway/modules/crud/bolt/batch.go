package bolt

import (
	"context"
	"fmt"

	"github.com/spaceuptech/space-cloud/gateway/model"
)

func (b *Bolt) Batch(ctx context.Context, project string, req *model.BatchRequest) ([]int64, error) {
	return nil, fmt.Errorf("batch operation not supported for selected database")
}
