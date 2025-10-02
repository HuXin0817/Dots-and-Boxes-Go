#include "BoardV1.h"

#include "../model/Box.h"
#include "../model/Edge.h"
#include "../model/EdgeBoxMapper.h"

int
BoardV1::Add(Edge edge) {
  BoardV0::Add(edge);
  return EdgeCountOfBox::Add(edge);
}

[[nodiscard]] Edge
BoardV1::FindNotContainsEdgeInBox(Box box) const {
  assert(EdgeCountOfBox::operator[](box) == 3);
  for (auto edge : EdgeBoxMapper::BoxNearEdges[box]) {
    if (NotContains(edge)) {
      return edge;
    }
  }
  assert(false);
  return -1;
}

[[nodiscard]] Edge
BoardV1::FindScoreableEdge() const {
  for (int box = 0; box < Box::Max; box++) {
    if (EdgeCountOfBox::operator[](box) == 3) {
      return FindNotContainsEdgeInBox(Box(box));
    }
  }
  return -1;
}

[[nodiscard]] BoardV1&
BoardV1::GetBoardV1() {
  return *this;
}

[[nodiscard]] const BoardV1&
BoardV1::GetBoardV1() const {
  return *this;
}
