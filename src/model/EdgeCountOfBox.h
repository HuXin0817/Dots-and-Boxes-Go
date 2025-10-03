#pragma once

#include "../common/Array.h"
#include "EdgeBoxMapper.h"

class EdgeCountOfBox : public Array<int, Box::Max> {
  public:
  EdgeCountOfBox() = default;

  int
  Add(Edge e) {
    int s = 0;
    for (auto box : EdgeBoxMapper::EdgeNearBoxes.At(e)) {
      At(box)++;
      assert(At(box) <= 4);
      if (At(box) == 4) {
        s++;
      }
    }
    return s;
  }

  int
  MaxCount(Edge e) const {
    int c = 0;
    for (auto box : EdgeBoxMapper::EdgeNearBoxes.At(e)) {
      c = std::max(c, At(box));
    }
    return c;
  }
};
