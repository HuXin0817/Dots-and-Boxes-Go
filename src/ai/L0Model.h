#pragma once

#include "../board/BoardV2.h"
#include "../common/List.h"
#include "../common/Span.h"
#include "Interface.h"

class L0Model final : public AIInterface {
  friend class L1Model;
  friend class L2Model;

  public:
  L0Model() = default;

  Span<Edge>
  BestCandidateEdges(const BoardV2& board) override {
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

  private:
  List<Edge, Edge::Max> EnemyUnscoreableEdges;
  List<Edge, Edge::Max> ScoreableEdges;
};