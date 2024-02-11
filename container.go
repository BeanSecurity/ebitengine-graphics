package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
)

type Container struct {
	Pos gmath.Pos

	Rotation *gmath.Rad

	objects []object

	visible  bool
	disposed bool
}

func NewContainer() *Container {
	return &Container{
		objects: make([]object, 0, 4),
	}
}

func (c *Container) IsDisposed() bool {
	return c.disposed
}

func (c *Container) Dispose() {
	for _, o := range c.objects {
		o.Dispose()
	}
	c.disposed = true
}

// IsVisible reports whether this container is visible.
// Use SetVisibility to change this flag value.
//
// If container is invisible, none of its objects will be rendered during the Draw call.
func (c *Container) IsVisible() bool {
	return c.visible
}

// SetVisibility changes the Visible flag value.
// It can be used to show or hide the container.
// Use IsVisible to get the current flag value.
func (c *Container) SetVisibility(visible bool) {
	c.visible = visible
}

func (c *Container) Draw(dst *ebiten.Image) {
	c.DrawWithOptions(dst, DrawOptions{})
}

func (c *Container) AddChild(o object) {
	c.objects = append(c.objects, o)
}

func (c *Container) DrawWithOptions(dst *ebiten.Image, opts DrawOptions) {
	if !c.visible {
		return
	}

	opts.Offset = opts.Offset.Add(c.Pos.Resolve())
	if c.Rotation != nil {
		opts.Rotation += *c.Rotation
	}

	liveObjects := c.objects[:0]
	for _, o := range c.objects {
		if o.IsDisposed() {
			continue
		}
		liveObjects = append(liveObjects, o)
		o.DrawWithOptions(dst, opts)
	}
	c.objects = liveObjects
}
