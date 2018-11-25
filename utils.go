package main

import (
	"encoding/binary"
	"github.com/PucklaMotzer09/gohomeengine/src/frameworks/GTK/gtk"
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"github.com/PucklaMotzer09/mathgl/mgl32"
	"math"
)

func uint32ToString(i uint32) string {
	var buffer [4]byte
	n := binary.PutUvarint(buffer[:], uint64(i))
	return string(buffer[:n])
}

func stringToUint32(str string) uint32 {
	i, _ := binary.Uvarint([]byte(str))
	return uint32(i)
}

func loadModel(name, fileContents, fileName string) {
	gohome.ErrorMgr.Log("Load", "Model", name)
	gohome.ResourceMgr.LoadLevelString(name, string(fileContents), fileName, true)

	level := gohome.ResourceMgr.GetLevel(name)
	if level != nil && len(level.LevelObjects) != 0 {
		for i := 0; i < len(level.LevelObjects); i++ {
			model := level.LevelObjects[i].Model3D
			if model != nil {
				if loaded_models == nil {
					loaded_models = make(map[uint32]*gohome.Model3D)
				}
				if placable_models == nil {
					placable_models = make(map[uint32]*PlaceableModel)
				}
				loaded_models[object_id] = model

				lbl := gtk.LabelNew(model.Name)
				lbl.ToGObject().SetData("ID", uint32ToString(object_id))
				lb_assets.Insert(lbl.ToWidget(), -1)
				lbl.ToWidget().Show()

				var pm PlaceableModel
				pm.Name = model.Name
				pm.Filename = fileName
				pm.ID = object_id
				placable_models[object_id] = &pm
				object_id++
			}
		}

	} else {
		gohome.ErrorMgr.Error("Load", "Model", "Loaded model is broken")
	}
}

func loadLoadableModels() {
	for i := 0; i < len(loadable_models); i++ {
		loadModel(loadable_models[i].Name, loadable_models[i].FileContents, loadable_models[i].Filename)
	}
	loadable_models = loadable_models[:0]
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
	smooth_deltas[current_smooth_delta][0] = dx
	smooth_deltas[current_smooth_delta][1] = dy
	dx, dy = smoothDeltas()
	if mgl32.Abs(dx) > MAX_DELTA || mgl32.Abs(dy) > MAX_DELTA {
		return
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
	smooth_zooms[current_smooth_zoom] = zoom
	zoom = smoothZooms()
	camera_zoom = mgl32.Clamp(camera_zoom-zoom, MIN_ZOOM, MAX_ZOOM)
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
	camera.Position = camera.Position.Add(vec)
	camera_center = camera.Position.Add(camera.LookDirection.Mul(camera_zoom))
}

func handleMoveArrowClick() {
	var m *PlacedModel
	var ok bool
	if m, ok = placed_models[place_id-1]; ok {
		transform_start_pos = m.Entity3D.Transform.GetPosition()
	} else {
		return
	}

	pointsx, pointsy, pointsz := arrows.GetMoveHitboxes()

	mpos := gohome.InputMgr.Mouse.ToScreenPosition()
	quadx, quady, quadz := gohome.QuadMath2D(pointsx), gohome.QuadMath2D(pointsy), gohome.QuadMath2D(pointsz)

	if quadx.IntersectsPoint(mpos) {
		arrows.IsTransforming = X_AXIS
		is_transforming = true
		arrows.CalculatePoints()
		transform_start = getAxisProjectedPos(mpos, X_AXIS, m)
	} else if quady.IntersectsPoint(mpos) {
		arrows.IsTransforming = Y_AXIS
		is_transforming = true
		arrows.CalculatePoints()
		transform_start = getAxisProjectedPos(mpos, Y_AXIS, m)
	} else if quadz.IntersectsPoint(mpos) {
		arrows.IsTransforming = Z_AXIS
		is_transforming = true
		arrows.CalculatePoints()
		transform_start = getAxisProjectedPos(mpos, Z_AXIS, m)
	}
}

func getAxisProjectedPos(screenPos mgl32.Vec2, axis uint8, m *PlacedModel) mgl32.Vec3 {
	var plane gohome.PlaneMath3D

	switch axis {
	case X_AXIS:
		plane = getBestPlane(X_PLANES)
	case Y_AXIS:
		plane = getBestPlane(Y_PLANES)
	case Z_AXIS:
		plane = getBestPlane(Z_PLANES)
	}

	changePlanePoint(&plane, m.Entity3D.Transform.Position)

	mray := gohome.ScreenPositionToRay(screenPos)
	planePos := mray.PlaneIntersect(camera.Position, plane.Normal, plane.Point)
	planePos = planePos.Project(arrows.points3D[0], arrows.points3D[1])

	return planePos
}

func handleTransforming() {
	if !is_transforming {
		return
	}
	m, ok := placed_models[place_id-1]
	if !ok {
		return
	}

	arrows.SetPosition()

	planePos := getAxisProjectedPos(gohome.InputMgr.Mouse.ToScreenPosition(), arrows.IsTransforming, m)
	if current_mode == MODE_MOVE {
		m.Entity3D.Transform.Position = planePos.Sub(transform_start).Add(transform_start_pos)
	}
}

func getBestPlane(planes [4]gohome.PlaneMath3D) gohome.PlaneMath3D {
	minDot := float32(math.Acos(float64(planes[0].Normal.Dot(camera.LookDirection))))
	minIndex := 0

	for i := 1; i < 4; i++ {
		dot := float32(math.Acos(float64(planes[i].Normal.Dot(camera.LookDirection))))
		if dot < minDot {
			minDot = dot
			minIndex = i
		}
	}

	return planes[minIndex]
}

func changePlanePoint(plane *gohome.PlaneMath3D, position mgl32.Vec3) {
	if plane.Normal.X() != 0.0 {
		plane.Point[0] = position.X()
	} else if plane.Normal.Y() != 0.0 {
		plane.Point[1] = position.Y()
	} else {
		plane.Point[2] = position.Z()
	}
}
