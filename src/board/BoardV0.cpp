#include "BoardV0.h"

BoardV0::BoardV0() {
  for (int i = 0; i < Edge::Max; i++) {
    EdgeIndexes[i] = i;
    Edges[i] = i;
  }
}

void
BoardV0::Add(Edge edge) {
  assert(NotContains(edge));
  int nowStep = NowStep();
  Edge nowEdge = Edges[nowStep];
  int edgeIndex = EdgeIndexes[edge];
  std::swap(Edges[edgeIndex], Edges[nowStep]);
  EdgeIndexes[edge] = nowStep;
  EdgeIndexes[nowEdge] = edgeIndex;
  Go();
}

bool
BoardV0::Contains(Edge e) const {
  return EdgeIndexes[e] < NowStep();
}

bool
BoardV0::NotContains(Edge e) const {
  return EdgeIndexes[e] >= NowStep();
}

[[nodiscard]] std::span<const Edge>
BoardV0::EmptyEdges() const {
  return {Edges.begin() + NowStep(), Edges.begin() + Edge::Max};
}

[[nodiscard]] std::span<const Edge>
BoardV0::MoveRecord() const {
  return {Edges.begin(), Edges.begin() + NowStep()};
}
