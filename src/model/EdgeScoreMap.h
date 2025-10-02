#pragma once

#include "EdgeList.h"

class EdgeScoreMap {
  public:
  void
  Reset();

  void
  Add(Edge edge, int score);

  EdgeScoreMap
  operator+=(const EdgeScoreMap& other);

  std::span<const Edge>
  Export();

  private:
  std::array<int, Edge::Max> Time{};
  std::array<int, Edge::Max> Score{};
  EdgeList BestEdges;
};