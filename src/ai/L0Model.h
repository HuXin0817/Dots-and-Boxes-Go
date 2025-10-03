#pragma once

#include "../common/List.h"
#include "../common/Span.h"
#include "Interface.h"

class L0Model final : public AIInterface {
  friend class L1Model;
  friend class L2Model;

  public:
  L0Model() = default;

  Span<const Edge>
  BestCandidateEdges(const BoardV2& board) override;

  private:
  List<Edge, Edge::Max> EnemyUnscoreableEdges;
  List<Edge, Edge::Max> ScoreableEdges;
};