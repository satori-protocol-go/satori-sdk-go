package client

const (
	// 单位: ms
	DEFAULT_TIMEOUT = 10000
)

type SignNo int32

const (
	EVENT    SignNo = 0
	PING     SignNo = 1
	PONG     SignNo = 2
	IDENTIFY SignNo = 3
	READY    SignNo = 4
)
