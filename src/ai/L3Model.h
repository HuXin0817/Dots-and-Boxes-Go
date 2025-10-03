#pragma once

#include "../board/BoardV2.h"
#include "../model/EdgeScoreMap.h"
#include "L2Model.h"

class L3Model final : public AIInterface {
  friend class L4Model;

  public:
  explicit L3Model(int SearchTime = 10000) : SearchTime(SearchTime) {
  }

  Span<Edge>
  BestCandidateEdges(const BoardV2& board) override {
    if (auto edges = SubModel.BestCandidateEdges(board); edges.Size() == 1) {
      return edges;
    }

    ScoreMap.Reset();
    int times = SearchTime / board.RemainStep() + 1;
    while (times--) {
      AuxBoard.Reset(board.GetBoardV1());
      Edge edge = RandomChoice(SubModel.BestCandidateEdges(AuxBoard));
      AuxBoard.Add(edge);
      while (AuxBoard.Gaming()) {
        AuxBoard.Add(RandomChoice(SubModel.BestCandidateEdges(AuxBoard)));
      }
      ScoreMap.Add(edge, AuxBoard.Score());
    }

    return ScoreMap.Export();
  }

  private:
  L2Model SubModel;
  int SearchTime;
  BoardV2 AuxBoard;
  EdgeScoreMap ScoreMap;
};
