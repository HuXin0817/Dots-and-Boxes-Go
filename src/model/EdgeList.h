#pragma once

#include <cassert>
#include <span>

#include "Edge.h"

class EdgeList {
  public:
  EdgeList() = default;

  void
  Reset(Edge e) {
    m[0] = e;
    len = 1;
  }

  void
  Clear() {
    len = 0;
  }

  [[nodiscard]] bool
  Empty() const {
    return len == 0;
  }

  void
  Append(Edge e) {
    assert(len < Edge::Max);
    m[len++] = e;
  }

  [[nodiscard]] std::span<const Edge>
  Export() const {
    return {m.begin(), m.begin() + len};
  }

  private:
  std::array<Edge, Edge::Max> m;
  int len = 0;
};
