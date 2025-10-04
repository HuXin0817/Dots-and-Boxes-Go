#pragma once

#include "../common/Config.h"

template <int size>
class Square {
  public:
  Square(int x, int y) : v(x * Size + y) {
  }

  int
  X() const {
    return v / size;
  }

  int
  Y() const {
    return v % size;
  }

  static constexpr int Size = size;
  static constexpr int Max = size * size;

  V(Square)
};

using Dot = Square<BoardSize + 1>;
using Box = Square<BoardSize>;
