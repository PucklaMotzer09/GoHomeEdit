package main

import (
	"github.com/PucklaMotzer09/GoHomeEngine/src/gohome"
	"github.com/PucklaMotzer09/mathgl/mgl32"
	"sync"
)

func (this *Arrows) initRotate() {
	/*this.rotateX.InitModel(gohome.ResourceMgr.GetLevel("Arrows").GetModel("Block_Cube"))
	this.rotateY.InitModel(gohome.ResourceMgr.GetLevel("Arrows").GetModel("Block_Cube").Copy())
	this.rotateZ.InitModel(gohome.ResourceMgr.GetLevel("Arrows").GetModel("Block_Cube").Copy())*/

	/*this.rotateX.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = X_COLOR
	this.rotateY.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = Y_COLOR
	this.rotateZ.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = Z_COLOR*/

	/*this.rotateX.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(180.0), mgl32.Vec3{0.0, 0.0, 1.0})
	this.rotateY.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(-90.0), mgl32.Vec3{0.0, 0.0, 1.0})
	this.rotateZ.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(90.0), mgl32.Vec3{0.0, 1.0, 0.0})*/

	/*this.rotateX.Transform.IgnoreParentRotation = true
	this.rotateX.Transform.IgnoreParentScale = true
	this.rotateX.DepthTesting = false
	this.rotateX.RenderLast = true
	this.rotateY.Transform.IgnoreParentRotation = true
	this.rotateY.Transform.IgnoreParentScale = true
	this.rotateY.DepthTesting = false
	this.rotateY.RenderLast = true
	this.rotateZ.Transform.IgnoreParentRotation = true
	this.rotateZ.Transform.IgnoreParentScale = true
	this.rotateZ.DepthTesting = false
	this.rotateZ.RenderLast = true*/

	/*gohome.RenderMgr.AddObject(&this.rotateX)
	gohome.RenderMgr.AddObject(&this.rotateY)
	gohome.RenderMgr.AddObject(&this.rotateZ)*/
}

func (this *Arrows) initScale() {
	this.scaleX.InitModel(gohome.ResourceMgr.GetLevel("Arrows").GetModel("Block_Cube.001"))
	this.scaleY.InitModel(gohome.ResourceMgr.GetLevel("Arrows").GetModel("Block_Cube.001").Copy())
	this.scaleZ.InitModel(gohome.ResourceMgr.GetLevel("Arrows").GetModel("Block_Cube.001").Copy())

	this.scaleX.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = X_COLOR
	this.scaleY.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = Y_COLOR
	this.scaleZ.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = Z_COLOR

	this.scaleX.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(180.0), mgl32.Vec3{0.0, 0.0, 1.0})
	this.scaleY.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(-90.0), mgl32.Vec3{0.0, 0.0, 1.0})
	this.scaleZ.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(90.0), mgl32.Vec3{0.0, 1.0, 0.0})

	this.scaleX.Transform.IgnoreParentRotation = true
	this.scaleX.Transform.IgnoreParentScale = true
	this.scaleX.DepthTesting = false
	this.scaleX.RenderLast = true
	this.scaleY.Transform.IgnoreParentRotation = true
	this.scaleY.Transform.IgnoreParentScale = true
	this.scaleY.DepthTesting = false
	this.scaleY.RenderLast = true
	this.scaleZ.Transform.IgnoreParentRotation = true
	this.scaleZ.Transform.IgnoreParentScale = true
	this.scaleZ.DepthTesting = false
	this.scaleZ.RenderLast = true

	gohome.RenderMgr.AddObject(&this.scaleX)
	gohome.RenderMgr.AddObject(&this.scaleY)
	gohome.RenderMgr.AddObject(&this.scaleZ)
}

func (this *Arrows) initMove() {
	this.translateX.InitModel(gohome.ResourceMgr.GetLevel("Arrows").GetModel("Arrow_Cone"))
	this.translateY.InitModel(gohome.ResourceMgr.GetLevel("Arrows").GetModel("Arrow_Cone").Copy())
	this.translateZ.InitModel(gohome.ResourceMgr.GetLevel("Arrows").GetModel("Arrow_Cone").Copy())

	this.translateX.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = X_COLOR
	this.translateY.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = Y_COLOR
	this.translateZ.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = Z_COLOR

	this.translateX.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(-90.0), mgl32.Vec3{0.0, 1.0, 0.0})
	this.translateY.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(90.0), mgl32.Vec3{1.0, 0.0, 0.0})
	this.translateZ.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(180.0), mgl32.Vec3{1.0, 0.0, 0.0})

	this.translateX.Transform.IgnoreParentRotation = true
	this.translateX.Transform.IgnoreParentScale = true
	this.translateX.DepthTesting = false
	this.translateX.RenderLast = true
	this.translateY.Transform.IgnoreParentRotation = true
	this.translateY.Transform.IgnoreParentScale = true
	this.translateY.DepthTesting = false
	this.translateY.RenderLast = true
	this.translateZ.Transform.IgnoreParentRotation = true
	this.translateZ.Transform.IgnoreParentScale = true
	this.translateZ.DepthTesting = false
	this.translateZ.RenderLast = true

	gohome.RenderMgr.AddObject(&this.translateX)
	gohome.RenderMgr.AddObject(&this.translateY)
	gohome.RenderMgr.AddObject(&this.translateZ)
}

func (this *Arrows) setScaleMove() {
	cam := camera.Position

	txs := this.translateX.Transform.GetPosition().Sub(cam).Len() * ARROWS_SIZE
	tys := this.translateY.Transform.GetPosition().Sub(cam).Len() * ARROWS_SIZE
	tzs := this.translateZ.Transform.GetPosition().Sub(cam).Len() * ARROWS_SIZE

	this.translateX.Transform.Scale = [3]float32{txs, txs, txs}
	this.translateY.Transform.Scale = [3]float32{tys, tys, tys}
	this.translateZ.Transform.Scale = [3]float32{tzs, tzs, tzs}
}

func (this *Arrows) setScaleScale() {
	cam := camera.Position

	sxs := this.scaleX.Transform.GetPosition().Sub(cam).Len() * ARROWS_SIZE
	sys := this.scaleY.Transform.GetPosition().Sub(cam).Len() * ARROWS_SIZE
	szs := this.scaleZ.Transform.GetPosition().Sub(cam).Len() * ARROWS_SIZE

	this.scaleX.Transform.Scale = [3]float32{sxs, sxs, sxs}
	this.scaleY.Transform.Scale = [3]float32{sys, sys, sys}
	this.scaleZ.Transform.Scale = [3]float32{szs, szs, szs}

}

func (this *Arrows) setScaleRotate() {
	/*cam := camera.Position

	rxs := this.rotateX.Transform.GetPosition().Sub(cam).Len() * ARROWS_SIZE
	rys := this.rotateY.Transform.GetPosition().Sub(cam).Len() * ARROWS_SIZE
	rzs := this.rotateZ.Transform.GetPosition().Sub(cam).Len() * ARROWS_SIZE

	this.rotateX.Transform.Scale = [3]float32{rxs, rxs, rxs}
	this.rotateY.Transform.Scale = [3]float32{rys, rys, rys}
	this.rotateZ.Transform.Scale = [3]float32{rzs, rzs, rzs}
	*/
}

func (this *Arrows) setVisibleMove() {
	this.translateX.Visible = true
	this.translateY.Visible = true
	this.translateZ.Visible = true
}

func (this *Arrows) setInvisibleMove() {
	this.translateX.Visible = false
	this.translateY.Visible = false
	this.translateZ.Visible = false
}

func (this *Arrows) setVisibleScale() {
	this.scaleX.Visible = true
	this.scaleY.Visible = true
	this.scaleZ.Visible = true
}

func (this *Arrows) setInvisibleScale() {
	this.scaleX.Visible = false
	this.scaleY.Visible = false
	this.scaleZ.Visible = false
}

func (this *Arrows) setVisibleRotate() {
	/*this.rotateX.Visible = true
	this.rotateY.Visible = true
	this.rotateZ.Visible = true*/
}

func (this *Arrows) setInvisibleRotate() {
	/*this.rotateX.Visible = false
	this.rotateY.Visible = false
	this.rotateZ.Visible = false*/
}

func (this *Arrows) calculateAllMatrices() {
	var wg sync.WaitGroup

	wg.Add(5)
	go func() {
		camera.CalculateViewMatrix()
		wg.Done()
	}()
	go func() {
		gohome.RenderMgr.Projection3D.CalculateProjectionMatrix()
		wg.Done()
	}()
	go func() {
		this.translateX.Transform.CalculateTransformMatrix(&gohome.RenderMgr, -1)
		wg.Done()
	}()
	go func() {
		this.translateY.Transform.CalculateTransformMatrix(&gohome.RenderMgr, -1)
		wg.Done()
	}()
	go func() {
		this.translateZ.Transform.CalculateTransformMatrix(&gohome.RenderMgr, -1)
		wg.Done()
	}()
	wg.Wait()
}

func convert3Dto2D(pos mgl32.Vec3, pos2 *mgl32.Vec2, wg *sync.WaitGroup) {
	vmat := camera.GetViewMatrix()
	pmat := gohome.RenderMgr.Projection3D.GetProjectionMatrix()
	mat := pmat.Mul4(vmat)
	pos4 := mat.Mul4x1(pos.Vec4(1))
	pos3 := pos4.Div(pos4.W()).Vec3()
	nres := gohome.Render.GetNativeResolution()

	*pos2 = pos3.Vec2()
	*pos2 = pos2.MulVec([2]float32{1.0, -1.0}).Add([2]float32{1.0, 1.0}).Div(2.0).MulVec(nres)
	wg.Done()
}

func (this *Arrows) centerArrows() {
	this.translateX.Transform.Position = [3]float32{0.0, 0.0, 0.0}
	this.translateY.Transform.Position = [3]float32{0.0, 0.0, 0.0}
	this.translateZ.Transform.Position = [3]float32{0.0, 0.0, 0.0}

	this.scaleX.Transform.Position = [3]float32{0.0, 0.0, 0.0}
	this.scaleY.Transform.Position = [3]float32{0.0, 0.0, 0.0}
	this.scaleZ.Transform.Position = [3]float32{0.0, 0.0, 0.0}

	/*this.rotateX.Transform.Position = [3]float32{0.0, 0.0, 0.0}
	this.rotateY.Transform.Position = [3]float32{0.0, 0.0, 0.0}
	this.rotateZ.Transform.Position = [3]float32{0.0, 0.0, 0.0}*/
}
