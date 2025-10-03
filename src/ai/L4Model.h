#pragma once

#include <omp.h>

#include <thread>

#include "Interface.h"
#include "L3Model.h"

class L4Model final : public AIInterface {
  public:
  static constexpr int SubModelSearchTime = 1000;

  explicit L4Model(int GroupNumber = 100) : GroupNumber(GroupNumber) {
  }

  Span<const Edge>
  BestCandidateEdges(const BoardV2& board) override;

  private:
  int GroupNumber;
};