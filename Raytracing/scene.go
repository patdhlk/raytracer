package scene

import (
	objects "de/vorlesung/projekt/raytracer/SceneObjects"
	"image"
	imageColor "image/color"
	"log"
	"math"
)

var global = 0.25

type Scene struct {
	eye      *objects.Vector
	grid     *Grid
	elements []SceneObject
	ambient  *objects.Vector
	skyColor *objects.Vector
	light    *Light
}

func NewScene(eye *objects.Vector, grid *Grid) *Scene {
	if grid.TopLeft().X() != grid.BottomRight().X() {
		panic("not same x values on corners")
	}
	tmp := new(Scene)
	tmp.SetEye(eye)
	tmp.SetGrid(grid)
	tmp.SetElements(make([]SceneObject, 0, 0))
	tmp.SetAmbient(objects.NewVector(0.0, 0.0, 0.0))
	tmp.SetSkyColor(objects.NewVector(0.0, 0.0, 0.0))
	return tmp
}

func (p *Scene) Render(x, y int, superSample int) image.Image {
	log.Println("Start rendering (function) Image\n")
	tmp_x := x * superSample
	tmp_y := y * superSample
	tmp_img := make([][]objects.Vector, tmp_x)

	rasterStart := p.grid.TopLeft()
	rasterSizeZ := math.Abs(rasterStart.Z()-p.grid.BottomRight().Z()) / float64(tmp_x)
	rasterSizeY := -math.Abs(rasterStart.Y()-p.grid.BottomRight().Y()) / float64(tmp_y)

	for i := 0; i < tmp_x; i++ {
		tmp_img[i] = make([]objects.Vector, tmp_y)
		for j := 0; j < tmp_y; j++ {

			posZ := rasterStart.Z() + (float64(i)+0.5)*rasterSizeZ
			posY := rasterStart.Y() + (float64(j)+0.5)*rasterSizeY
			gridPos := objects.NewVector(rasterStart.X(), posY, posZ)

			var color, _ = p.followRay(p.eye, gridPos.Sub(p.eye), nil, 8)
			if color == nil {
				color = p.skyColor
			}

			tmp_img[i][j] = *color
		}
	}
	return scale(tmp_img, tmp_x, tmp_y, superSample)
}

func scale(in [][]objects.Vector, size_x, size_y, factor int) *image.RGBA {
	var out = image.NewRGBA(image.Rect(0, 0, size_x/factor, size_y/factor))

	for i := 0; i < size_x; i += factor {
		for j := 0; j < size_y; j += factor {
			var tmpVal = objects.NewVector(0.0, 0.0, 0.0)
			for k := 0; k < factor; k++ {
				for l := 0; l < factor; l++ {
					tmpVal = tmpVal.Add(&in[i+k][j+l])
				}
			}

			tmpVal = tmpVal.DivVal(float64(factor * factor))
			var outColor = new(imageColor.RGBA)
			outColor.R = uint8(tmpVal.X() * 255)
			outColor.G = uint8(tmpVal.Y() * 255)
			outColor.B = uint8(tmpVal.Z() * 255)
			outColor.A = 255
			out.Set(i/factor, j/factor, outColor)
		}
	}
	return out
}

func (p *Scene) followRay(position, direction *objects.Vector, ignored SceneObject, depthLeft uint8) (*objects.Vector, *objects.Vector) {
	var Ray = objects.NewRay(position, direction)

	var intersectObject, intersectPos, color, normal, diffuse, specularIntensity, specularPower, reflectivity, nearestDist = p.intersectAll(Ray, ignored)

	if math.IsInf(nearestDist, 1) {
		color = nil
	} else {
		var phongColor *objects.Vector
		var phongSpecular *objects.Vector
		if p.light != nil {
			phongColor, phongSpecular = p.calcPhong(intersectObject, intersectPos, Ray, normal, diffuse, specularIntensity/5, specularPower*2)
		} else {
			phongColor = objects.NewVector(1.0, 1.0, 1.0)
			phongSpecular = objects.NewVector(0.0, 0.0, 0.0)
		}

		color = color.Mul(phongColor).Add(phongSpecular).Limit(0, 1)

		if depthLeft > 0 && reflectivity > 0 {
			var reflectColor, reflectPos = p.followRay(intersectPos, direction.Reflect(normal), intersectObject, depthLeft-1)
			if reflectColor != nil {
				color = color.MulVal(1-reflectivity).Add(reflectColor.MulVal(reflectivity)).Limit(0, 1)

				// Ambient occlusion
				var reflectDistance = reflectPos.Sub(intersectPos).Length()
				if reflectDistance < global {
					color = color.Mul(objects.NewVector(1.0, 1.0, 1.0).MulVal(reflectDistance/global).Limit(0.25, 1))
				}
			} else {
				color = color.MulVal(1-reflectivity).Add(p.skyColor.MulVal(reflectivity)).Limit(0, 1)
			}
		}
	}
	return color, intersectPos
}

func (p *Scene) checkInShadow(intersectPos, directionToLight *objects.Vector, intersectObject SceneObject) bool {
	r := objects.NewRay(intersectPos, directionToLight)
	_, pos, _, _, _, _, _, _, _ := p.intersectAll(r, intersectObject)
	return pos != nil
}

func (p *Scene) calcPhong(intersectObject SceneObject, intersectPos *objects.Vector, ray *objects.Ray, normal *objects.Vector,
	diffuse float64, specularIntensity float64,
	specularPower float64) (phongColor, phongSpecular *objects.Vector) {

	phongColor = p.ambient
	phongSpecular = objects.NewVector(0.0, 0.0, 0.0)

	directionToLight := p.light.Position().Sub(intersectPos).Normalized()
	if !p.checkInShadow(intersectPos, directionToLight, intersectObject) {

		directionFromLight := directionToLight.MulVal(-1)
		normal = normal.Normalized()
		phongDiffuse := p.light.Color().MulVal(diffuse*directionToLight.Dot(normal)).Limit(0, 1)

		directionLightOut := directionFromLight.Reflect(normal).Normalized()
		directionIntersectEye := ray.Origin().Sub(intersectPos).Normalized()
		specularAmount := directionLightOut.Dot(directionIntersectEye)
		if specularAmount < 0 {
			specularAmount = 0
		}
		phongSpecular = p.light.Color().MulVal(specularIntensity*math.Pow(specularAmount, specularPower)).Limit(0, 1)

		phongColor = phongColor.Add(phongDiffuse).Limit(0, 1)
	}
	return
}

func (p *Scene) intersectAll(Ray *objects.Ray, ignored SceneObject) (intersectObject SceneObject,
	intersectPos, color, normal *objects.Vector,
	diffuse, specularIntensity, specularPower,
	reflectivity, nearestDist float64) {

	nearestDist = math.Inf(1)
	for _, element := range p.elements {
		if ignored != nil && element == ignored {
			continue
		}

		var intersectP, col, norm, dif, spInt, spPow, ref = element.Intersection(Ray)
		if intersectP == nil {
			continue
		}
		var length = intersectP.Sub(Ray.Origin()).Length()
		if length < nearestDist {
			intersectObject = element
			intersectPos = intersectP
			color = col
			normal = norm
			diffuse = dif
			specularIntensity = spInt
			specularPower = spPow
			reflectivity = ref
			nearestDist = length
		}
	}
	return
}

func (p *Scene) Eye() *objects.Vector                 { return p.eye }
func (p *Scene) Grid() *Grid                          { return p.grid }
func (p *Scene) Elements() []SceneObject              { return p.elements }
func (p *Scene) Ambient() *objects.Vector             { return p.ambient }
func (p *Scene) Light() *Light                        { return p.light }
func (p *Scene) SkyColor() *objects.Vector            { return p.skyColor }
func (p *Scene) SetEye(eye *objects.Vector)           { p.eye = eye }
func (p *Scene) SetGrid(grid *Grid)                   { p.grid = grid }
func (p *Scene) SetElements(elements []SceneObject)   { p.elements = elements }
func (p *Scene) SetAmbient(ambient *objects.Vector)   { p.ambient = ambient }
func (p *Scene) SetLight(light *Light)                { p.light = light }
func (p *Scene) SetSkyColor(skyColor *objects.Vector) { p.skyColor = skyColor }
func (p *Scene) AddElement(element SceneObject)       { p.elements = append(p.elements, element) }
