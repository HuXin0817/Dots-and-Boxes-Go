#include "EdgeScoreMap.h"

void
EdgeScoreMap::Reset() {
  Time = std::array<int, Edge::Max>{};
  Score = std::array<int, Edge::Max>{};
  BestEdges.Clear();
}

void
EdgeScoreMap::Add(Edge edge, int score) {
  Time[edge]++;
  Score[edge] += score;
}

EdgeScoreMap
EdgeScoreMap::operator+=(const EdgeScoreMap& other) {
  for (int i = 0; i < Edge::Max; i++) {
    Time[i] += other.Time[i];
    Score[i] += other.Score[i];
  }
  return *this;
}

std::span<const Edge>
EdgeScoreMap::Export() {
  float maxs = 0.0;
  for (int e = 0; e < Edge::Max; e++) {
    if (Time[e] > 0) {
      float s = static_cast<float>(Score[e]) / static_cast<float>(Time[e]);
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
