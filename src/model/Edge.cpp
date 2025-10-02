#include "Edge.h"

Edge::Edge(int v) : v(v) {
}

Edge::Edge(Dot dot1, Dot dot2) : v(std::get<0>(DotMapper)[dot1][dot2]) {
}

[[nodiscard]] Dot
Edge::dot1() const {
  return std::get<1>(DotMapper)[v];
}

[[nodiscard]] Dot
Edge::dot2() const {
  return std::get<2>(DotMapper)[v];
}

Edge::
operator int() const {
  return v;
}

inline std::tuple<std::array<std::array<int, Dot::Max>, Dot::Max>,
                  std::array<Dot, Edge::Max>,
                  std::array<Dot, Edge::Max>>
    Edge::DotMapper = [] {
      std::array<std::array<int, Dot::Max>, Dot::Max> DotsToEdges{};
      std::array<Dot, Max> EdgeToDot1{};
      std::array<Dot, Max> EdgeToDot2{};

      int e = 0;
      for (int x = 0; x < Dot::Size; x++) {
        for (int y = 0; y < Dot::Size; y++) {
          Dot d1(x, y);
          if (int x1 = x + 1; x1 < Dot::Size) {
            Dot d2(x1, y);
            DotsToEdges[d1][d2] = e;
            EdgeToDot1[e] = d1;
            EdgeToDot2[e] = d2;
            e++;
          }
          if (int y1 = y + 1; y1 < Dot::Size) {
            Dot d2(x, y1);
            DotsToEdges[d1][d2] = e;
            EdgeToDot1[e] = d1;
            EdgeToDot2[e] = d2;
            e++;
          }
        }
      }

      return std::make_tuple(DotsToEdges, EdgeToDot1, EdgeToDot2);
    }();
