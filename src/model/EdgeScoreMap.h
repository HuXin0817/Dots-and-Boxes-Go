#pragma once

#include "../common/Array.h"
#include "../common/List.h"
#include "../common/Span.h"
#include "Edge.h"

class EdgeScoreMap {
  public:
  void
  Reset();

  void
  Add(Edge edge, int score);

  EdgeScoreMap
  operator+=(const EdgeScoreMap& other);

  Span<const Edge>
  Export();

  private:
  Array<int, Edge::Max> Time{};
  Array<int, Edge::Max> Score{};
  List<Edge, Edge::Max> BestEdges;
};