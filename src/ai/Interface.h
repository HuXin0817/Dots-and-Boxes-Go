#pragma once

#include <vector>

#include "../board/BoardV2.h"
#include "../model/Edge.h"

class AIInterface {
  public:
  virtual ~AIInterface() = default;

  virtual std::span<const Edge>
  BestCandidateEdges(const BoardV2& board) = 0;
};