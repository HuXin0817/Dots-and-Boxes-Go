#pragma once

#include <cassert>
#include <span>

#include "../model/Step.h"

class BoardV0 : public Step {
  public:
  BoardV0();

  void
  Add(Edge edge);

  bool
  Contains(Edge e) const;

  bool
  NotContains(Edge e) const;

  [[nodiscard]] std::span<const Edge>
  EmptyEdges() const;

  [[nodiscard]] std::span<const Edge>
  MoveRecord() const;

  private:
  std::array<Edge, Edge::Max> Edges;
  std::array<int, Edge::Max> EdgeIndexes{};
};