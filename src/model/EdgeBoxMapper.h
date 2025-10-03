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
