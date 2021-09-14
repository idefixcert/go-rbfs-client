package ping

import (
	"context"
	"net/http"

	"github.com/rtbrick/go-rbfs-client/pkg/rbfs/state"
	"github.com/stretchr/testify/mock"
)

// ensure, that mockActionsAPI does implement ActionsAPI.
var _ ActionsAPI = &mockActionsAPI{}

type mockActionsAPI struct {
	mock.Mock
}

func (m *mockActionsAPI) ActionsPingPost(ctx context.Context, destinationIP string,
	localVarOptionals *state.ActionsApiActionsPingPostOpts) (state.PingStatus, *http.Response, error) {
	args := m.Called(ctx, destinationIP, localVarOptionals)
	status, ok := args.Get(0).(state.PingStatus)
	if !ok {
		status = state.PingStatus{}
	}
	return status, nil, args.Error(1)
}
