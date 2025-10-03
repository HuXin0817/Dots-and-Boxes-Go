#pragma once

#include "../common/Config.h"

class Dot {
  public:
  Dot() = default;

  Dot(int v) : v(v) {
  }

  Dot(int x, int y) : v(x * Size + y) {
  }

  [[nodiscard]] int
  X() const {
    return v / Size;
  }

  [[nodiscard]] int
  Y() const {
    return v % Size;
  }

  static constexpr int Size = BoardSize + 1;

  static constexpr int Max = Size * Size;

  operator int() const {
    return v;
  }

  private:
  int v = 0;
};
