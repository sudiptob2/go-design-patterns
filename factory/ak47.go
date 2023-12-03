package main

type AK47 struct {
	Name  string
	Speed int
}

func (ak *AK47) Power() float32 {
	return float32(ak.Speed) * 0.4
}
func (ak *AK47) GetName() string {
	return ak.Name
}

func newAK47() IGun {
	return &AK47{
		Name:  "Ak 47 gun",
		Speed: 4,
	}
}
