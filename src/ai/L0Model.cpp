#include "L0Model.h"

#include "../board/BoardV2.h"

std::span<const Edge>
L0Model::BestCandidateEdges(const BoardV2& board) {
  EnemyUnscoreableEdges.Clear();
  ScoreableEdges.Clear();
  const auto EmptyEdges = board.EmptyEdges();

  for (auto edge : EmptyEdges) {
    if (int maxCount = board.MaxCount(edge); maxCount == 3) {
      ScoreableEdges.Append(edge);
    } else if (maxCount < 2) {
      EnemyUnscoreableEdges.Append(edge);
    }
  }

  if (!ScoreableEdges.Empty()) {
    return ScoreableEdges.Export();
  }
  if (!EnemyUnscoreableEdges.Empty()) {
    return EnemyUnscoreableEdges.Export();
  }

  return {EmptyEdges.begin(), EmptyEdges.end()};
}
