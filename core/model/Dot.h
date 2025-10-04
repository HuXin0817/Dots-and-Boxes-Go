#pragma once

#include "../common/Config.h"

class Dot {
  public:
  Dot(int x, int y) : v(x * Size + y) {
  }

  int
  X() const {
    return v / Size;
  }

  int
  Y() const {
    return v % Size;
  }

  static constexpr int Size = BoardSize + 1;
  static constexpr int Max = Size * Size;

  V(Dot)
};
