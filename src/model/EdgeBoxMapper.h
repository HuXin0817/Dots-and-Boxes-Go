#pragma once

#include "../common/Array.h"
#include "Box.h"
#include "Edge.h"

class EdgeBoxMapper {
  friend class Box;
  friend class Edge;

  public:
  static Array<Array<Edge, 4>, Box::Max> BoxNearEdges;
  static Array<std::vector<Box>, Edge::Max> EdgeNearBoxes;
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

inline Array<std::vector<Box>, Edge::Max> EdgeBoxMapper::EdgeNearBoxes = [] {
  auto nearBoxes = [](Edge e) {
    std::vector<Box> boxes;
    int x = e.dot2().X() - 1;
    int y = e.dot2().Y() - 1;
    if (x >= 0 && y >= 0) {
      boxes.emplace_back(x, y);
    }

    x = e.dot1().X();
    y = e.dot1().Y();
    if (x < BoardSize && y < BoardSize) {
      boxes.emplace_back(x, y);
    }
    return boxes;
  };

  Array<std::vector<Box>, Edge::Max> mapper;
  for (int e = 0; e < Edge::Max; e++) {
    mapper.At(e) = nearBoxes(Edge(e));
  }
  return mapper;
}();
