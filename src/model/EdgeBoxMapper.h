#pragma once

#include "../common/Array.h"
#include "../common/List.h"
#include "../common/Span.h"
#include "Box.h"
#include "Edge.h"

class EdgeBoxMapper {
  friend class Box;
  friend class Edge;

  public:
  static Array<Array<Edge, 4>, Box::Max> BoxNearEdges;
  static Array<Span<Box>, Edge::Max> EdgeNearBoxes;
};

inline Array<Array<Edge, 4>, Box::Max> EdgeBoxMapper::BoxNearEdges = [] {
  [[maybe_unused]] Edge dummy(0);

  Array<Array<Edge, 4>, Box::Max> BoxNearEdges;

  for (int i = 0; i < Box::Max; i++) {
    int x = i / BoardSize;
    int y = i % BoardSize;
    Dot d00(x, y);
    Dot d01(x + 1, y);
    Dot d10(x, y + 1);
    Dot d11(x + 1, y + 1);
    BoxNearEdges.At(i) = {
        Edge(d00, d01),
        Edge(d00, d10),
        Edge(d10, d11),
        Edge(d01, d11),
    };
  }

  return BoxNearEdges;
}();

inline Array<Span<Box>, Edge::Max> EdgeBoxMapper::EdgeNearBoxes = [] {
  static Array<Box, 2 * Edge::Max - BoardSize * 4> NearBoxes;
  int index = 0;

  auto nearBoxes = [&index](Edge e) -> Span<Box> {
    auto start = NearBoxes.begin() + index;

    int x = e.dot2().X() - 1;
    int y = e.dot2().Y() - 1;
    if (x >= 0 && y >= 0) {
      NearBoxes.At(index++) = Box(x, y);
    }

    x = e.dot1().X();
    y = e.dot1().Y();
    if (x < BoardSize && y < BoardSize) {
      NearBoxes.At(index++) = Box(x, y);
    }

    auto end = NearBoxes.begin() + index;

    return {start, end};
  };

  Array<Span<Box>, Edge::Max> mapper;
  for (int e = 0; e < Edge::Max; e++) {
    mapper.At(e) = nearBoxes(Edge(e));
  }
  return mapper;
}();
