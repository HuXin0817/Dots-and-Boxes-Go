#include "EdgeCountOfBox.h"

int
EdgeCountOfBox::Add(Edge e) {
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
EdgeCountOfBox::MaxCount(Edge e) const {
  int c = 0;
  for (auto box : EdgeBoxMapper::EdgeNearBoxes.At(e)) {
    c = std::max(c, At(box));
  }
  return c;
}
