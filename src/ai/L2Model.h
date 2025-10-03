#pragma once

#include "../common/List.h"
#include "../common/Span.h"
#include "L1Model.h"

class L2Model final : public AIInterface {
  public:
  L2Model() = default;

  Span<const Edge>
  BestCandidateEdges(const BoardV2& board) override;

  private:
  L1Model SubModel;
  BoardV2 AuxBoard;
  List<Edge, Edge::Max> SearchEdges;
};