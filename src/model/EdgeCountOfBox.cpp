#include "EdgeCountOfBox.h"

int
EdgeCountOfBox::Add(Edge e) {
  int s = 0;
  for (auto box : EdgeBoxMapper::EdgeNearBoxes[e]) {
    at(box)++;
    assert(at(box) <= 4);
    if (operator[](box) == 4) {
      s++;
    }
  }
  return s;
}

int
EdgeCountOfBox::MaxCount(Edge e) const {
  int c = 0;
  for (auto box : EdgeBoxMapper::EdgeNearBoxes[e]) {
    c = std::max(c, operator[](box));
  }
  return c;
}
