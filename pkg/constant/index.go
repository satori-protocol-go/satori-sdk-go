package constant

const (
	// 单位: ms
	DEFAULT_TIMEOUT = 10000
)

type SignNum uint64

const (
	SIGN_NUM_EVENT    SignNum = 0
	SIGN_NUM_PING     SignNum = 1
	SIGN_NUM_PONG     SignNum = 2
	SIGN_NUM_IDENTIFY SignNum = 3
	SIGN_NUM_READY    SignNum = 4
)
