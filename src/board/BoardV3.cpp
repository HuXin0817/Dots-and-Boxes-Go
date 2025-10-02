#include "BoardV3.h"

void
BoardV3::Reset(const BoardV1& nb) {
  GetBoardV1() = nb;
  ScoreableEdges.Clear();
}

int
BoardV3::Add(Edge edge) {
  int score = BoardV1::Add(edge);
  for (auto box : EdgeBoxMapper::EdgeNearBoxes[edge]) {
    if (EdgeCountOfBox::operator[](box) == 3) {
      Edge edgeToAdd = FindNotContainsEdgeInBox(box);
      ScoreableEdges.Append(edgeToAdd);
    }
  }
  return score;
}

int
BoardV3::MaxObtainableScore(int minScore) {
  int score = 0;
  while (Gaming()) {
    if (ScoreableEdges.Empty()) {
      if (Edge e = FindScoreableEdge(); static_cast<int>(e) != -1) {
        ScoreableEdges.Append(e);
      } else {
        break;
      }
    }
    Edge edge = ScoreableEdges.Pop();
    if (Contains(edge)) {
      continue;
    }
    int addScore = Add(edge);
    assert(addScore > 0);
    score += addScore;
    if (score >= minScore) {
      break;
    }
  }
  return score;
}

[[nodiscard]] bool
BoardV3::ScoreableEdgesEmpty() const {
  return ScoreableEdges.Empty();
}
