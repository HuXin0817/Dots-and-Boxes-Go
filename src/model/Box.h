#pragma once

#include "Dot.h"

class Box {
  public:
  Box() = default;

  Box(int v) : v(v) {
  }

  Box(int x, int y) : v(x * BoardSize + y) {
  }

  static constexpr int Size = BoardSize;
  static constexpr int Max = Size * Size;

  operator int() const {
    return v;
  }

  private:
  int v = 0;
};
