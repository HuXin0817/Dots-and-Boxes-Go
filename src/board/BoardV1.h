#pragma once

#include "../model/Box.h"
#include "../model/Edge.h"
#include "../model/EdgeCountOfBox.h"
#include "BoardV0.h"

class BoardV1 : public BoardV0, public EdgeCountOfBox {
  public:
  BoardV1() = default;

  int
  Add(Edge edge);

  [[nodiscard]] Edge
  FindNotContainsEdgeInBox(Box box) const;

  [[nodiscard]] Edge
  FindScoreableEdge() const;

  [[nodiscard]] BoardV1&
  GetBoardV1();

  [[nodiscard]] const BoardV1&
  GetBoardV1() const;
};
