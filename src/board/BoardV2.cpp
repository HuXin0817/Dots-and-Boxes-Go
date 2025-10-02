#include "BoardV2.h"

void
BoardV2::Reset(const BoardV1& nb) {
  GetBoardV1() = nb;
  ScoreMap::Reset();
}

int
BoardV2::Add(Edge e) {
  int score = BoardV1::Add(e);
  ScoreMap::Add(score);
  return score;
}

bool
BoardV2::Gaming() const {
  return ScoreMap::Gaming() && BoardV1::Gaming();
}
