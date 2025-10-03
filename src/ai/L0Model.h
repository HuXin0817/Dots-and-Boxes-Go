#pragma once

#include "Interface.h"

class L0Model final : public AIInterface {
  friend class L1Model;
  friend class L2Model;

  public:
  L0Model() = default;

  std::span<const Edge>
  BestCandidateEdges(const BoardV2& board) override;

  private:
  List<Edge, Edge::Max> EnemyUnscoreableEdges;
  List<Edge, Edge::Max> ScoreableEdges;
};