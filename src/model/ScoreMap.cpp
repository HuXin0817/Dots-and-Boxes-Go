#include "ScoreMap.h"

ScoreMap::ScoreMap() {
  Reset();
}

void
ScoreMap::Reset() {
  Player1Score = 0;
  Player2Score = 0;
  Turn = Player1Turn;
}

void
ScoreMap::Add(int s) {
  if (s == 0) {
    Turn = !Turn;
    return;
  }
  if (Turn == Player1Turn) {
    Player1Score += s;
  } else {
    Player2Score += s;
  }
}

[[nodiscard]] int
ScoreMap::Score() const {
  return Player1Score - Player2Score;
}

[[nodiscard]] int
ScoreMap::GetScore(int player) const {
  return (player == 0) ? Player1Score : Player2Score;
}

[[nodiscard]] bool
ScoreMap::Gaming() const {
  return Player1Score < MinWinnerScore && Player2Score < MinWinnerScore;
}
