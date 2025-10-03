#pragma once

#include "../common/Array.h"
#include "../common/Span.h"
#include "../model/Edge.h"
#include "../model/Step.h"

class BoardV0 : public Step {
  public:
  BoardV0() {
    for (int i = 0; i < Edge::Max; i++) {
      EdgeIndexes.At(i) = i;
      Edges.At(i) = i;
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
  Contains(Edge e) const {
    return EdgeIndexes.At(e) < NowStep();
  }

  bool
  NotContains(Edge e) const {
    return EdgeIndexes.At(e) >= NowStep();
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