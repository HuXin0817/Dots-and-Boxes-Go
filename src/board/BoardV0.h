#pragma once

#include "../common/Array.h"
#include "../common/Span.h"
#include "../model/Edge.h"
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

  [[nodiscard]] Span<const Edge>
  EmptyEdges() const;

  [[nodiscard]] Span<const Edge>
  MoveRecord() const;

  private:
  Array<Edge, Edge::Max> Edges;
  Array<int, Edge::Max> EdgeIndexes{};
};