#pragma once

#include "Dot.h"

class Box {
  public:
  Box(int x, int y) : v(x * BoardSize + y) {
  }

  static constexpr int Size = BoardSize;
  static constexpr int Max = Size * Size;

  V(Box)
};
