#pragma once

#include <tuple>

#include "../common/Array.h"
#include "Square.h"

class Edge {
  public:
  static constexpr int Max = 2 * BoardSize * (BoardSize + 1);

  Edge(Dot dot1, Dot dot2) : v(std::get<0>(DotMapper).At(dot1).At(dot2)) {
  }

  Dot
  Dot1() const {
    return std::get<1>(DotMapper).At(v);
  }

  Dot
  Dot2() const {
    return std::get<2>(DotMapper).At(v);
  }

  bool
  Rotate() {
    return Dot1().Y() == Dot2().Y();
  }

  V(Edge)

  private:
  static std::tuple<Array<Array<int, Dot::Max>, Dot::Max>, Array<Dot, Max>, Array<Dot, Max>>
  GetDotMapper() {
    Array<Array<int, Dot::Max>, Dot::Max> DotsToEdges{};
    Array<Dot, Max> EdgeToDot1{};
    Array<Dot, Max> EdgeToDot2{};

    int edge = 0;
    for (int x = 0; x < Dot::Size; x++) {
      for (int y = 0; y < Dot::Size; y++) {
        Dot d1(x, y);
        if (int x1 = x + 1; x1 < Dot::Size) {
          Dot d2(x1, y);
          DotsToEdges.At(d1).At(d2) = edge;
          EdgeToDot1.At(edge) = d1;
          EdgeToDot2.At(edge) = d2;
          edge++;
        }
        if (int y1 = y + 1; y1 < Dot::Size) {
          Dot d2(x, y1);
          DotsToEdges.At(d1).At(d2) = edge;
          EdgeToDot1.At(edge) = d1;
          EdgeToDot2.At(edge) = d2;
          edge++;
        }
      }
    }

    return std::make_tuple(DotsToEdges, EdgeToDot1, EdgeToDot2);
  }

  static inline std::tuple<Array<Array<int, Dot::Max>, Dot::Max>, Array<Dot, Max>, Array<Dot, Max>>
      DotMapper = GetDotMapper();
};
