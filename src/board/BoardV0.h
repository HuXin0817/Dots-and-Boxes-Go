#pragma once

#include "../common/Array.h"
#include "../common/Span.h"
#include "../model/Edge.h"
#include "../model/Step.h"

class BoardV0 : public Step {
  public:
  BoardV0() {
    for (int edge = 0; edge < Edge::Max; edge++) {
      EdgeIndexes.At(edge) = edge;
      Edges.At(edge) = edge;
    }
  }

  void
  Add(Edge edge) {
    assert(NotContains(edge));
    int nowStep = NowStep();
    Edge nowEdge = Edges.At(nowStep);
    int edgeIndex = EdgeIndexes.At(edge);
    std::swap(Edges.At(edgeIndex), Edges.At(nowStep));
    EdgeIndexes.At(edge) = nowStep;
    EdgeIndexes.At(nowEdge) = edgeIndex;
    Go();
  }

  bool
  Contains(Edge edge) const {
    return EdgeIndexes.At(edge) < NowStep();
  }

  bool
  NotContains(Edge edge) const {
    return EdgeIndexes.At(edge) >= NowStep();
  }

  Span<Edge>
  EmptyEdges() const {
    return {Edges.begin() + NowStep(), Edges.begin() + Edge::Max};
  }

  Span<Edge>
  MoveRecord() const {
    return {Edges.begin(), Edges.begin() + NowStep()};
  }

  private:
  Array<Edge, Edge::Max> Edges;
  Array<int, Edge::Max> EdgeIndexes{};
};
