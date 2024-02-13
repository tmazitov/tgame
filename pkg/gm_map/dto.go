package gm_map

import "github.com/hajimehoshi/ebiten/v2"

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

type IMapObj interface {
	Draw(screen *ebiten.Image, camera *Camera)
}

const CameraBorderSize float64 = 64
