#pragma once

#include "../board/BoardV2.h"
#include "../model/EdgeScoreMap.h"
#include "L2Model.h"

class L3Model final : public AIInterface {
  friend class L4Model;

  public:
  explicit L3Model(int SearchTime = 10000) : SearchTime(SearchTime) {
  }

  Span<const Edge>
  BestCandidateEdges(const BoardV2& board) override {
    if (auto l = SubModel.BestCandidateEdges(board); l.size() == 1) {
      return l;
    }

    ScoreMap.Reset();
    int Times = SearchTime / board.RemainStep() + 1;
    while (Times--) {
      AuxBoard.Reset(board.GetBoardV1());
      auto e = RandomChoice(SubModel.BestCandidateEdges(AuxBoard));
      AuxBoard.Add(e);
      while (AuxBoard.Gaming()) {
        AuxBoard.Add(RandomChoice(SubModel.BestCandidateEdges(AuxBoard)));
      }
      ScoreMap.Add(e, AuxBoard.Score());
    }

    return ScoreMap.Export();
  }

  private:
  L2Model SubModel;
  int SearchTime;
  BoardV2 AuxBoard;
  EdgeScoreMap ScoreMap;
};