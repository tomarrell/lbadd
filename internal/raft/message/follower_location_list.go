package message

//go:generate protoc --go_out=. follower_location_list.proto

var _ Message = (*FollowerLocationListRequest)(nil)
var _ Message = (*FollowerLocationListResponse)(nil)

// NewFollowerLocationListRequest creates a new follower-location-list-request
// message with the given parameters.
func NewFollowerLocationListRequest() *FollowerLocationListRequest {
	return &FollowerLocationListRequest{}
}

// Kind returns KindFollowerLocationListRequest.
func (*FollowerLocationListRequest) Kind() Kind {
	return KindFollowerLocationListRequest
}

// NewFollowerLocationListResponse creates a new follower-location-list-response
// message with the given parameters.
func NewFollowerLocationListResponse(followerLocations []string) *FollowerLocationListResponse {
	return &FollowerLocationListResponse{
		FollowerAddress: followerLocations,
	}
}

// Kind returns KindFollowerLocationListResponse.
func (*FollowerLocationListResponse) Kind() Kind {
	return KindFollowerLocationListResponse
}
