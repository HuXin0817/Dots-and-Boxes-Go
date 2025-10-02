#pragma once

#include "Box.h"
#include "Edge.h"

class EdgeBoxMapper {
  friend class Box;
  friend class Edge;

  public:
  static std::array<std::array<Edge, 4>, Box::Max> BoxNearEdges;
  static std::array<std::vector<Box>, Edge::Max> EdgeNearBoxes;
};
