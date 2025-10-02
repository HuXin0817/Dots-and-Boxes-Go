#pragma once

#include "../board/BoardV2.h"
#include "../model/EdgeScoreMap.h"
#include "L2Model.h"

class L3Model final : public AIInterface {
  friend class L4Model;

  public:
  explicit L3Model(int SearchTime = 10000) : SearchTime(SearchTime) {
  }

  std::span<const Edge>
  BestCandidateEdges(const BoardV2& board) override;

  private:
  L2Model SubModel;
  int SearchTime;
  BoardV2 AuxBoard;
  EdgeScoreMap ScoreMap;
};