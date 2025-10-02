#pragma once

#include <tuple>
#include <vector>

#include "Dot.h"

class Edge {
  public:
  static constexpr int Max = 2 * BoardSize * (BoardSize + 1);

  Edge() = default;

  Edge(int v);

  Edge(Dot dot1, Dot dot2);

  [[nodiscard]] Dot
  dot1() const;

  [[nodiscard]] Dot
  dot2() const;

  operator int() const;

  private:
  int v = 0;

  static std::tuple<std::array<std::array<int, Dot::Max>, Dot::Max>,
                    std::array<Dot, Max>,
                    std::array<Dot, Max>>
      DotMapper;
};
