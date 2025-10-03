#pragma once

#include "../model/Box.h"
#include "../model/Edge.h"
#include "../model/EdgeCountOfBox.h"
#include "BoardV0.h"

class BoardV1 : public BoardV0, public EdgeCountOfBox {
  public:
  BoardV1() = default;

  int
  Add(Edge edge) {
    BoardV0::Add(edge);
    return EdgeCountOfBox::Add(edge);
  }

  [[nodiscard]] Edge
  FindNotContainsEdgeInBox(Box box) const {
    assert(EdgeCountOfBox::At(box) == 3);
    for (auto edge : EdgeBoxMapper::BoxNearEdges.At(box)) {
      if (NotContains(edge)) {
        return edge;
      }
    }
    assert(false);
    return -1;
  }

  [[nodiscard]] Edge
  FindScoreableEdge() const {
    for (int box = 0; box < Box::Max; box++) {
      if (EdgeCountOfBox::At(box) == 3) {
        return FindNotContainsEdgeInBox(Box(box));
      }
    }
    return -1;
  }

  [[nodiscard]] BoardV1&
  GetBoardV1() {
    return *this;
  }

  [[nodiscard]] const BoardV1&
  GetBoardV1() const {
    return *this;
  }
};
