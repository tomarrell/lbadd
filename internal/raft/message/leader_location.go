package message

//go:generate protoc --go_out=. leader_location.proto

var _ Message = (*LeaderLocationRequest)(nil)
var _ Message = (*LeaderLocationResponse)(nil)

// NewLeaderLocationRequest creates a new leader-location-request message with
// the given parameters.
func NewLeaderLocationRequest() *LeaderLocationRequest {
	return &LeaderLocationRequest{}
}

// Kind returns KindLeaderLocationRequest.
func (*LeaderLocationRequest) Kind() Kind {
	return KindLeaderLocationRequest
}

// NewLeaderLocationResponse creates a new leader-location-response message with
// the given parameters.
func NewLeaderLocationResponse(leaderAddress string) *LeaderLocationResponse {
	return &LeaderLocationResponse{
		LeaderAddress: leaderAddress,
	}
}

// Kind returns KindLeaderLocationResponse.
func (*LeaderLocationResponse) Kind() Kind {
	return KindLeaderLocationResponse
}
