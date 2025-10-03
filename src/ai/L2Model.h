#pragma once

#include "L1Model.h"

class L2Model final : public AIInterface {
  public:
  L2Model() = default;

  std::span<const Edge>
  BestCandidateEdges(const BoardV2& board) override;

  private:
  L1Model SubModel;
  BoardV2 AuxBoard;
  List<Edge, Edge::Max> SearchEdges;
};