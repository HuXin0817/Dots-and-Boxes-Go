#pragma once

#include "../board/BoardV3.h"
#include "../common/Span.h"
#include "Interface.h"
#include "L0Model.h"

class L1Model final : public AIInterface {
  friend class L2Model;

  public:
  Span<const Edge>
  BestCandidateEdges(const BoardV2& board) override {
    if (auto l = SubModel.BestCandidateEdges(board);
        !SubModel.EnemyUnscoreableEdges.Empty() || !SubModel.ScoreableEdges.Empty()) {
      return l;
    }

    int mins = Box::Max + 1;
    auto& Candidate = SubModel.EnemyUnscoreableEdges;
    assert(Candidate.Empty());

    for (auto e : board.EmptyEdges()) {
      AuxBoard.Reset(board.GetBoardV1());
      int v = AuxBoard.Add(e);
      assert(v == 0);

      int s = AuxBoard.MaxObtainableScore(mins);
      if (s < mins) {
        mins = s;
        Candidate.Reset(e);
      } else if (s == mins) {
        Candidate.Append(e);
      }
    }

    return Candidate.Export();
  }

  private:
  L0Model SubModel;
  BoardV3 AuxBoard;
};
