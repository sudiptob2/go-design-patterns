package main

type Rifle struct {
	Name  string
	Speed int
}

func (rf *Rifle) Power() float32 {
	return float32(rf.Speed) * 0.6
}
func (rf *Rifle) GetName() string {
	return rf.Name
}

func newRifle() IGun {
	return &Rifle{
		Name:  "Rifle gun",
		Speed: 5,
	}
}
