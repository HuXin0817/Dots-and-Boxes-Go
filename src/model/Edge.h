#pragma once

#include <tuple>
#include <vector>

#include "Dot.h"

class Edge {
  public:
  static constexpr int Max = 2 * BoardSize * (BoardSize + 1);

  Edge() = default;

  Edge(int v) : v(v) {
  }

  Edge(Dot dot1, Dot dot2) : v(std::get<0>(DotMapper)[dot1][dot2]) {
  }

  [[nodiscard]] Dot
  dot1() const {
    return std::get<1>(DotMapper)[v];
  }

  [[nodiscard]] Dot
  dot2() const {
    return std::get<2>(DotMapper)[v];
  }

  operator int() const {
    return v;
  }

  private:
  int v = 0;

  static std::tuple<std::array<std::array<int, Dot::Max>, Dot::Max>,
                    std::array<Dot, Max>,
                    std::array<Dot, Max>>
      DotMapper;
};
