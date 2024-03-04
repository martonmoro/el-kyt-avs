package types

import (
	"errors"

	cstaskmanager "github.com/martonmoro/el-kyt-avs/contracts/bindings/KYTTaskManager"
)

type TaskResponseData struct {
	TaskResponse              cstaskmanager.IKYTTaskManagerTaskResponse
	TaskResponseMetadata      cstaskmanager.IKYTTaskManagerTaskResponseMetadata
	NonSigningOperatorPubKeys []cstaskmanager.BN254G1Point
}

var (
	NoErrorInTaskResponse = errors.New("100. Task response is valid")
)
