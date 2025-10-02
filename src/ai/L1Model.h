#pragma once

#include "../board/BoardV3.h"
#include "Interface.h"
#include "L0Model.h"

class L1Model final : public AIInterface {
  friend class L2Model;

  public:
  std::span<const Edge>
  BestCandidateEdges(const BoardV2& board) override;

  private:
  L0Model SubModel;
  BoardV3 AuxBoard;
};
