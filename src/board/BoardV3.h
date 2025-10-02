#pragma once


#include "../model/Edge.h"
#include "../model/EdgeQueue.h"
#include "BoardV1.h"

class BoardV3 : public BoardV1 {
  public:
  BoardV3() = default;

  void
  Reset(const BoardV1& nb);

  int
  Add(Edge edge);

  int
  MaxObtainableScore(int minScore);

  [[nodiscard]] bool
  ScoreableEdgesEmpty() const;

  private:
  EdgeQueue ScoreableEdges;
};
