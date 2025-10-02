#pragma once

#include "../defs.h"

class Dot {
  public:
  Dot() = default;

  Dot(int v);

  Dot(int x, int y);

  [[nodiscard]] int
  X() const;

  [[nodiscard]] int
  Y() const;

  static constexpr int Size = BoardSize + 1;

  static constexpr int Max = Size * Size;

  operator int() const;

  private:
  int v = 0;
};
