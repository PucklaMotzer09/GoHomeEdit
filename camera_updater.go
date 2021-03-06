package main

import (
	"github.com/PucklaMotzer09/GoHomeEngine/src/gohome"
	"github.com/PucklaMotzer09/mathgl/mgl32"
)

type CameraUpdater struct {
}

func (this *CameraUpdater) Init() {
	camera.Init()
	camera.LookAt(mgl32.Vec3{MID_ZOOM, MID_ZOOM, MID_ZOOM}, camera_center, mgl32.Vec3{0.0, 1.0, 0.0})
	gohome.RenderMgr.SetCamera3D(&camera, 0)

	gohome.UpdateMgr.AddObject(this)
}

func (this *CameraUpdater) Update(delta_time float32) {
	updateCamera()
}

func updateCamera() {
	updateCameraZoom()
	updateCameraRotation()
	updateCameraPanning()
}

var smooth_deltas [NUM_SMOOTH_DELTAS][2]float32
var current_smooth_deltas int = NUM_SMOOTH_DELTAS
var current_smooth_delta int = 0

func smoothDeltas() (dx float32, dy float32) {
	var sumx, sumy float32 = 0.0, 0.0
	for i := 0; i < current_smooth_deltas; i++ {
		sumx += smooth_deltas[i][0]
		sumy += smooth_deltas[i][1]
	}
	sumx /= float32(current_smooth_deltas)
	sumy /= float32(current_smooth_deltas)
	dx = sumx
	dy = sumy

	current_smooth_deltas++
	if current_smooth_deltas > NUM_SMOOTH_DELTAS {
		current_smooth_deltas = NUM_SMOOTH_DELTAS
	}
	current_smooth_delta++
	if current_smooth_delta == NUM_SMOOTH_DELTAS {
		current_smooth_delta = 0
	}

	return
}

func resetSmoothDeltas() {
	for i := 0; i < NUM_SMOOTH_DELTAS; i++ {
		smooth_deltas[i][0] = 0.0
		smooth_deltas[i][1] = 0.0
	}
	current_smooth_deltas = NUM_SMOOTH_DELTAS
	current_smooth_delta = 0
}

func updateCameraRotation() {
	if is_transforming {
		return
	}
	dx, dy := float32(gohome.InputMgr.Mouse.DPos[0]), float32(gohome.InputMgr.Mouse.DPos[1])
	if !gohome.InputMgr.IsPressed(gohome.MouseButtonRight) {
		dx, dy = 0.0, 0.0
	}
	if gohome.InputMgr.IsPressed(gohome.MouseButtonRight) && (gohome.InputMgr.IsPressed(gohome.KeyLeftShift) || gohome.InputMgr.IsPressed(gohome.KeyRightShift)) {
		dx, dy = 0.0, 0.0
		camera_pitch, camera_yaw = 3.1415/4.0, 3.1415/4.0
		gohome.RenderMgr.ReRender = true
	} else {
		smooth_deltas[current_smooth_delta][0] = dx
		smooth_deltas[current_smooth_delta][1] = dy
		dx, dy = smoothDeltas()
		if mgl32.Abs(dx) > MAX_DELTA || mgl32.Abs(dy) > MAX_DELTA {
			return
		}
	}
	if dx != 0.0 || dy != 0.0 {
		gohome.RenderMgr.ReRender = true
	}
	yaw, pitch := mgl32.DegToRad(-dx*CAM_ROTATE_VELOCITY), mgl32.DegToRad(dy*CAM_ROTATE_VELOCITY)

	if camera_pitch+pitch > mgl32.DegToRad(88.0) || camera_pitch+pitch < mgl32.DegToRad(-85.0) {
		pitch = 0.0
	}

	pos := mgl32.Vec3{0.0, 0.0, 1.0}
	look := mgl32.Vec3{0.0, 0.0, -1.0}
	up := mgl32.Vec3{0.0, 1.0, 0.0}
	relVec := pos

	rotateAxisPitch := up.Cross(look).Normalize()
	rotatePitch := mgl32.HomogRotate3D(camera_pitch, rotateAxisPitch)

	rotateAxisYaw := mgl32.Vec3{0.0, 1.0, 0.0}
	rotateYaw := mgl32.HomogRotate3D(camera_yaw, rotateAxisYaw)

	rotate := rotateYaw.Mul4(rotatePitch)

	relVec = rotate.Mul4x1(relVec.Vec4(0.0)).Vec3()

	camera.Position = camera_center.Add(relVec.Mul(camera_zoom))
	camera.LookDirection = camera_center.Sub(camera.Position).Normalize()

	camera_yaw += yaw
	camera_pitch += pitch
}

var smooth_zooms [NUM_SMOOTH_ZOOM]float32
var current_smooth_zoom int = 0

func smoothZooms() float32 {
	var sum float32 = 0.0
	for i := 0; i < NUM_SMOOTH_ZOOM; i++ {
		sum += smooth_zooms[i]
	}
	current_smooth_zoom++
	if current_smooth_zoom == NUM_SMOOTH_ZOOM {
		current_smooth_zoom = 0
	}
	return sum / float32(NUM_SMOOTH_ZOOM)
}

func updateCameraZoom() {
	if is_transforming {
		return
	}

	wy := float32(gohome.InputMgr.Mouse.Wheel[1])
	zoom := wy * CAM_ZOOM_VELOCITY
	if zoom != 0.0 && (gohome.InputMgr.IsPressed(gohome.KeyLeftShift) || gohome.InputMgr.IsPressed(gohome.KeyRightShift)) {
		camera_zoom = MID_ZOOM
		gohome.RenderMgr.ReRender = true
	} else {
		smooth_zooms[current_smooth_zoom] = zoom
		zoom = smoothZooms()
		camera_zoom = mgl32.Clamp(camera_zoom-zoom, MIN_ZOOM, MAX_ZOOM)
	}
	if zoom != 0.0 {
		gohome.RenderMgr.ReRender = true
	}

	PLACE_PLANE_DIST = camera_zoom
}

var smooth_pans [NUM_SMOOTH_PAN][2]float32
var current_smooth_pan int = 0

func smoothPans() (float32, float32) {
	var sumx, sumy float32 = 0.0, 0.0
	for i := 0; i < NUM_SMOOTH_PAN; i++ {
		sumx += smooth_pans[i][0]
		sumy += smooth_pans[i][1]
	}

	current_smooth_pan++
	if current_smooth_pan == NUM_SMOOTH_PAN {
		current_smooth_pan = 0
	}

	return sumx / float32(NUM_SMOOTH_PAN), sumy / float32(NUM_SMOOTH_PAN)
}

func updateCameraPanning() {
	if is_transforming {
		return
	}

	dx, dy := float32(gohome.InputMgr.Mouse.DPos[0]), float32(gohome.InputMgr.Mouse.DPos[1])
	if !gohome.InputMgr.IsPressed(gohome.MouseButtonMiddle) {
		dx, dy = 0.0, 0.0
	}
	smooth_pans[current_smooth_pan][0] = dx
	smooth_pans[current_smooth_pan][1] = dy
	dx, dy = smoothPans()
	if mgl32.Abs(dx) > MAX_DELTA || mgl32.Abs(dy) > MAX_DELTA {
		return
	}
	if dx != 0.0 || dy != 0.0 {
		gohome.RenderMgr.ReRender = true
	}

	panx, pany := dx*CAM_PAN_VELOCITY, dy*CAM_PAN_VELOCITY

	up := mgl32.Vec3{0.0, 1.0, 0.0}
	look := mgl32.Vec3{0.0, 0.0, -1.0}

	rotateAxisPitch := up.Cross(look).Normalize()
	rotatePitch := mgl32.HomogRotate3D(camera_pitch, rotateAxisPitch)
	rotateAxisYaw := mgl32.Vec3{0.0, 1.0, 0.0}
	rotateYaw := mgl32.HomogRotate3D(camera_yaw, rotateAxisYaw)

	up = rotateYaw.Mul4(rotatePitch).Mul4x1(up.Vec4(0.0)).Vec3()
	right := up.Cross(camera.LookDirection).Normalize()
	vec := up.Mul(pany).Add(right.Mul(panx))
	if gohome.InputMgr.IsPressed(gohome.MouseButtonMiddle) && (gohome.InputMgr.IsPressed(gohome.KeyLeftShift) || gohome.InputMgr.IsPressed(gohome.KeyRightShift)) {
		camera_center = mgl32.Vec3{0.0, 0.0, 0.0}
		camera.Position = camera_center.Add(camera.LookDirection.Mul(-camera_zoom))
	} else {
		camera.Position = camera.Position.Add(vec)
		camera_center = camera.Position.Add(camera.LookDirection.Mul(camera_zoom))
	}
}
