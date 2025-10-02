#pragma once

#include "../model/ScoreMap.h"
#include "BoardV1.h"

class BoardV2 : public BoardV1, public ScoreMap {
  public:
  BoardV2() = default;

  void
  Reset(const BoardV1& nb);

  int
  Add(Edge e);

  bool
  Gaming() const;
};