package patterns

type RemoteControl struct{
	vol volume
	o on
	mu mute
}
func (rc *RemoteControl) IncreaseVolume(){
	rc.vol.increase()
}
func (rc *RemoteControl) DecreaseVolume(){
	rc.vol.decrease()
}
func (rc *RemoteControl) MuteUnmute(){
	rc.mu.toggle()
}
func (rc *RemoteControl) OnOff(){
	rc.o.toggle()
}

type volume struct{
	Vol int
}
func (v *volume) increase(){
	v.Vol++
}
func (v *volume) decrease(){
	v.Vol--
}

type mute struct{
	mu bool
}
func (m *mute) toggle(){
	m.mu = !m.mu
}

type on struct{
	o bool
}
func (onoff *on) toggle(){
	onoff.o = !onoff.o
}


