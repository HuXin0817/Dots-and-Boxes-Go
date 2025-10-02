#pragma once

#include "Box.h"

static constexpr bool Player1Turn = true;
static constexpr bool Player2Turn = !Player1Turn;

static int MinWinnerScore = Box::Max / 2 + 1;

class ScoreMap {
  public:
  ScoreMap();

  void
  Reset();

  void
  Add(int s);

  [[nodiscard]] int
  Score() const;

  [[nodiscard]] int
  GetScore(int player) const;

  [[nodiscard]] bool
  Gaming() const;

  int Player1Score = 0;
  int Player2Score = 0;
  bool Turn = Player1Turn;
};