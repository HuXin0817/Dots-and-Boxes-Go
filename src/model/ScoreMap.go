package model

const (
	Player1Turn = 1
	Player2Turn = -Player1Turn

	minWinnerScore = int(MaxBox/2 + 1)
)

type ScoreMap struct {
	Turn         int
	Player1Score int
	Player2Score int
}

func (m *ScoreMap) Reset() {
	m.Turn = Player1Turn
	m.Player1Score = 0
	m.Player2Score = 0
}

func (m *ScoreMap) Add(s int) {
	if s == 0 {
		m.Turn = -m.Turn
		return
	}
	if m.Turn == Player1Turn {
		m.Player1Score += s
	} else {
		m.Player2Score += s
	}
}

func (m *ScoreMap) Score() int { return m.Player1Score - m.Player2Score }

func (m *ScoreMap) NotOver() bool {
	if m.Player1Score >= minWinnerScore {
		return false
	}
	if m.Player2Score >= minWinnerScore {
		return false
	}
	return true
}
