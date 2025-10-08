#pragma once

#include "Square.h"

class Edge {
  public:
  static constexpr int Max = 2 * BoardSize * (BoardSize + 1);

  Edge(Dot dot1, Dot dot2) {
    if (dot2 - dot1 == 1) {
      v = 2 * (dot1 - dot1 / (BoardSize + 1)) + 1;
    } else {
      v = 2 * dot1;
    }

    assert(Dot1() == dot1);
    assert(Dot2() == dot2);
  }

  Dot
  Dot1() const {
    return (v >> 1) + (v & 1) * (v >> 1) / BoardSize;
  }

  Dot
  Dot2() const {
    return (v >> 1) + (v & 1 ? (v >> 1) / BoardSize + 1 : BoardSize + 1);
  }

  bool
  Rotate() const {
    return v & 1;
  }

  std::string
  ToString() const {
    std::stringstream ss;
    ss << Dot1().ToString() << " -> " << Dot2().ToString();
    return ss.str();
  }

  V(Edge)
};
