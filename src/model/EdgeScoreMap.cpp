#include "EdgeScoreMap.h"

void
EdgeScoreMap::Reset() {
  Time = Array<int, Edge::Max>{};
  Score = Array<int, Edge::Max>{};
  BestEdges.Clear();
}

void
EdgeScoreMap::Add(Edge edge, int score) {
  Time.At(edge)++;
  Score.At(edge) += score;
}

EdgeScoreMap
EdgeScoreMap::operator+=(const EdgeScoreMap& other) {
  for (int i = 0; i < Edge::Max; i++) {
    Time.At(i) += other.Time.At(i);
    Score.At(i) += other.Score.At(i);
  }
  return *this;
}

Span<const Edge>
EdgeScoreMap::Export() {
  float maxs = 0.0;
  for (int e = 0; e < Edge::Max; e++) {
    if (Time.At(e) > 0) {
      float s = static_cast<float>(Score.At(e)) / static_cast<float>(Time.At(e));
      if (s > maxs || BestEdges.Empty()) {
        maxs = s;
        BestEdges.Reset(e);
      } else if (s == maxs) {
        BestEdges.Append(e);
      }
    }
  }
  return BestEdges.Export();
}
