#pragma once

#include "Dot.h"

class Box {
  public:
  Box() = default;

  Box(int v);

  Box(int x, int y);

  static constexpr int Size = BoardSize;
  static constexpr int Max = Size * Size;

  operator int() const;

  private:
  int v = 0;
};
