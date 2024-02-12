package gm_map

type MapLevel int

const (
	MapGroundLevel MapLevel = 1
)

type CameraArea int

const (
	NoneCameraArea        CameraArea = 0
	FreeCameraArea        CameraArea = 1
	TopBorderCameraArea   CameraArea = 2
	RightBorderCameraArea CameraArea = 3
	LeftBorderCameraArea  CameraArea = 4
	BotBorderCameraArea   CameraArea = 5

	TopLeftCornerArea  CameraArea = 6
	TopRightCornerArea CameraArea = 7
	BotLeftCornerArea  CameraArea = 8
	BotRightCornerArea CameraArea = 9
)

const CameraBorderSize float64 = 64
