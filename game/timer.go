package game

type Timer struct {
	currenTicks int
	targetTicks int
}

func NewTimer(targetTicks int) *Timer {
	return &Timer{
		currenTicks: 0,
		targetTicks: targetTicks,
	}
}

func (t *Timer) Update() {
	if t.currenTicks < t.targetTicks {
		t.currenTicks++
	}
}

func (t *Timer) IsReady() bool {
	return t.currenTicks >= t.targetTicks
}

func (t *Timer) Reset() {
	t.currenTicks = 0
}
