#pragma once

#include "../board/BoardV3.h"
#include "../common/Span.h"
#include "Interface.h"
#include "L0Model.h"

class L1Model final : public AIInterface {
  friend class L2Model;

  public:
  Span<Edge>
  BestCandidateEdges(const BoardV2& board) override {
    if (auto edges = SubModel.BestCandidateEdges(board);
        !SubModel.EnemyUnscoreableEdges.Empty() || !SubModel.ScoreableEdges.Empty()) {
      return edges;
    }

    int minScore = Box::Max + 1;
    auto& candidateEdges = SubModel.EnemyUnscoreableEdges;
    assert(candidateEdges.Empty());

    for (Edge edge : board.EmptyEdges()) {
      AuxBoard.Reset(board.GetBoardV1());
#ifdef NDEBUG
      AuxBoard.Add(edge);
#else
      assert(AuxBoard.Add(edge) == 0);
#endif
      int score = AuxBoard.MaxObtainableScore(minScore);
      if (score < minScore) {
        minScore = score;
        candidateEdges.Reset(edge);
      } else if (score == minScore) {
        candidateEdges.Append(edge);
      }
    }

    return candidateEdges.Export();
  }

  private:
  L0Model SubModel;
  BoardV3 AuxBoard;
};
