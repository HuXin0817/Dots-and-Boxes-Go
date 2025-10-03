#pragma once

#include "../board/BoardV2.h"
#include "../common/Span.h"
#include "../model/Edge.h"

class AIInterface {
  public:
  virtual ~AIInterface() = default;

  virtual Span<const Edge>
  BestCandidateEdges(const BoardV2& board) = 0;
};