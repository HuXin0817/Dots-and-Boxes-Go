#pragma once

#include <vector>

#include "../model/EdgeList.h"
#include "Interface.h"

class L0Model final : public AIInterface {
  friend class L1Model;
  friend class L2Model;

  public:
  L0Model() = default;

  std::span<const Edge>
  BestCandidateEdges(const BoardV2& board) override;

  private:
  EdgeList EnemyUnscoreableEdges;
  EdgeList ScoreableEdges;
};