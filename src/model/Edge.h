#pragma once

#include <tuple>

#include "../common/Array.h"
#include "Dot.h"

class Edge {
  public:
  static constexpr int Max = 2 * BoardSize * (BoardSize + 1);

  Edge() = default;

  Edge(int v) : v(v) {
  }

  Edge(Dot dot1, Dot dot2) : v(std::get<0>(DotMapper).At(dot1).At(dot2)) {
  }

  [[nodiscard]] Dot
  dot1() const {
    return std::get<1>(DotMapper).At(v);
  }

  [[nodiscard]] Dot
  dot2() const {
    return std::get<2>(DotMapper).At(v);
  }

  operator int() const {
    return v;
  }

  private:
  int v = 0;

  static std::tuple<Array<Array<int, Dot::Max>, Dot::Max>, Array<Dot, Max>, Array<Dot, Max>>
      DotMapper;
};

inline std::
    tuple<Array<Array<int, Dot::Max>, Dot::Max>, Array<Dot, Edge::Max>, Array<Dot, Edge::Max>>
        Edge::DotMapper = [] {
          Array<Array<int, Dot::Max>, Dot::Max> DotsToEdges{};
          Array<Dot, Max> EdgeToDot1{};
          Array<Dot, Max> EdgeToDot2{};

          int e = 0;
          for (int x = 0; x < Dot::Size; x++) {
            for (int y = 0; y < Dot::Size; y++) {
              Dot d1(x, y);
              if (int x1 = x + 1; x1 < Dot::Size) {
                Dot d2(x1, y);
                DotsToEdges.At(d1).At(d2) = e;
                EdgeToDot1.At(e) = d1;
                EdgeToDot2.At(e) = d2;
                e++;
              }
              if (int y1 = y + 1; y1 < Dot::Size) {
                Dot d2(x, y1);
                DotsToEdges.At(d1).At(d2) = e;
                EdgeToDot1.At(e) = d1;
                EdgeToDot2.At(e) = d2;
                e++;
              }
            }
          }

          return std::make_tuple(DotsToEdges, EdgeToDot1, EdgeToDot2);
        }();
