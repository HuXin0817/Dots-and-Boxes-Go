#pragma once

#include "Edge.h"

class Step {
  public:
  [[nodiscard]] bool
  Gaming() const {
    return v < Edge::Max;
  }

  [[nodiscard]] int
  RemainStep() const {
    return Edge::Max - v;
  }

  [[nodiscard]] int
  NowStep() const {
    return v;
  }

  void
  Go() {
    v++;
  }

  private:
  int v = 0;
};