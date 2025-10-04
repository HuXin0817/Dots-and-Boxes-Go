#pragma once

#include "../common/Config.h"

template <int A>
class Square {
  public:
  Square(int x, int y) : v(x * Size + y) {
  }

  int
  X() const {
    return v / Size;
  }

  int
  Y() const {
    return v % Size;
  }

  static constexpr int Size = A;
  static constexpr int Max = Size * Size;

  V(Square)
};

using Dot = Square<BoardSize + 1>;
using Box = Square<BoardSize>;
