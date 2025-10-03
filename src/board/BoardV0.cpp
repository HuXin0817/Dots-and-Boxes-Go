#include "BoardV0.h"

BoardV0::BoardV0() {
  for (int i = 0; i < Edge::Max; i++) {
    EdgeIndexes.At(i) = i;
    Edges.At(i) = i;
  }
}

void
BoardV0::Add(Edge edge) {
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
BoardV0::Contains(Edge e) const {
  return EdgeIndexes.At(e) < NowStep();
}

bool
BoardV0::NotContains(Edge e) const {
  return EdgeIndexes.At(e) >= NowStep();
}

[[nodiscard]] Span<const Edge>
BoardV0::EmptyEdges() const {
  return {Edges.begin() + NowStep(), Edges.begin() + Edge::Max};
}

[[nodiscard]] Span<const Edge>
BoardV0::MoveRecord() const {
  return {Edges.begin(), Edges.begin() + NowStep()};
}
