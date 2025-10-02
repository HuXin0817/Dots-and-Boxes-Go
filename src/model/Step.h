#pragma once

#include "Edge.h"

class Step {
  public:
  [[nodiscard]] bool
  Gaming() const;

  [[nodiscard]] int
  RemainStep() const;

  [[nodiscard]] int
  NowStep() const;

  void
  Go();

  private:
  int v = 0;
};