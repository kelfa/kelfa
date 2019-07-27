package session

func (ss *Sessions) Bounces() int {
	var bounches int
	for _, s := range ss.Sessions {
		if len(s.DataPoints) == 1 {
			bounches++
		}
	}
	return bounches
}

func (ss *Sessions) BounceRate() float64 {
	return float64(len(ss.Sessions)) / float64(ss.Bounces())
}
