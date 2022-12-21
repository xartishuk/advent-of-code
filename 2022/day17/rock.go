package main

type RockDispenser struct {
	i int
}

func (d *RockDispenser) Next(towerHeight int) Rock {
	d.i = (d.i + 1) % 5

	switch d.i {
	case 1:
		return &Plank{BaseRock{
			x: 2,
			y: towerHeight + 4,
			points: []point{
				{0, 0},
				{1, 0},
				{2, 0},
				{3, 0},
			},
		}}
	case 2:
		return &Plus{BaseRock{
			x: 2,
			y: towerHeight + 4,
			points: []point{
				{1, 0},
				{0, 1},
				{1, 1},
				{2, 1},
				{1, 2},
			},
		}}
	case 3:
		return &L{BaseRock{
			x: 2,
			y: towerHeight + 4,
			points: []point{
				{0, 0},
				{1, 0},
				{2, 0},
				{2, 1},
				{2, 2},
			},
		}}
	case 4:
		return &Rod{BaseRock{
			x: 2,
			y: towerHeight + 4,
			points: []point{
				{0, 0},
				{0, 1},
				{0, 2},
				{0, 3},
			},
		}}
	case 0:
		return &Box{BaseRock{
			x: 2,
			y: towerHeight + 4,
			points: []point{
				{0, 0},
				{1, 0},
				{0, 1},
				{1, 1},
			},
		}}
	default:
		panic("unreachable")
	}
}

type Rock interface {
	FallColliders() []point
	Fall()
	RightColliders() []point
	MoveRight()
	LeftColliders() []point
	MoveLeft()
	All() []point
}

type Plank struct{ BaseRock }

func (r *Plank) FallColliders() []point {
	return absPos(r.x, r.y, r.points)
}

func (r *Plank) RightColliders() []point {
	return absPos(r.x, r.y, r.points[3:])
}

func (r *Plank) LeftColliders() []point {
	return absPos(r.x, r.y, r.points[:1])
}

type Plus struct{ BaseRock }

func (r *Plus) FallColliders() []point {
	return absPos(r.x, r.y, []point{r.points[0], r.points[1], r.points[3]})
}

func (r *Plus) RightColliders() []point {
	return absPos(r.x, r.y, []point{r.points[0], r.points[3], r.points[4]})
}

func (r *Plus) LeftColliders() []point {
	return absPos(r.x, r.y, []point{r.points[0], r.points[1], r.points[4]})
}

type L struct{ BaseRock }

func (r *L) FallColliders() []point {
	return absPos(r.x, r.y, r.points[:3])
}

func (r *L) RightColliders() []point {
	return absPos(r.x, r.y, r.points[2:])
}

func (r *L) LeftColliders() []point {
	return absPos(r.x, r.y, []point{r.points[0], r.points[3], r.points[4]})
}

type Rod struct{ BaseRock }

func (r *Rod) FallColliders() []point {
	return absPos(r.x, r.y, r.points[:1])
}

func (r *Rod) RightColliders() []point {
	return absPos(r.x, r.y, r.points)
}

func (r *Rod) LeftColliders() []point {
	return absPos(r.x, r.y, r.points)
}

type Box struct{ BaseRock }

func (r *Box) FallColliders() []point {
	return absPos(r.x, r.y, r.points[:2])
}

func (r *Box) RightColliders() []point {
	return absPos(r.x, r.y, []point{r.points[1], r.points[3]})
}

func (r *Box) LeftColliders() []point {
	return absPos(r.x, r.y, []point{r.points[0], r.points[2]})
}

type BaseRock struct {
	x, y   int
	points []point
}

func (r *BaseRock) Fall() {
	r.y--
}

func (r *BaseRock) MoveRight() {
	r.x++
}

func (r *BaseRock) MoveLeft() {
	r.x--
}

func (r *BaseRock) All() []point {
	return absPos(r.x, r.y, r.points)
}

type point struct {
	x, y int
}

func absPos(x, y int, points []point) []point {
	absolute := make([]point, 0, len(points))
	for _, p := range points {
		absolute = append(absolute, point{
			x: x + p.x,
			y: y + p.y,
		})
	}

	return absolute
}
