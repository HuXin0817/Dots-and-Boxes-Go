#pragma once

#include "../model/ScoreMap.h"
#include "BoardV1.h"

class BoardV2 : public BoardV1, public ScoreMap {
  public:
  BoardV2() = default;

  void
  Reset(const BoardV1& nb) {
    GetBoardV1() = nb;
    ScoreMap::Reset();
  }

  int
  Add(Edge e) {
    int score = BoardV1::Add(e);
    ScoreMap::Add(score);
    return score;
  }

  bool
  Gaming() const {
    return ScoreMap::Gaming() && BoardV1::Gaming();
  }
};