#pragma once

#include <cassert>
#include <span>

#include "Edge.h"

class EdgeList {
  public:
  EdgeList() = default;

  void
  Reset(Edge e);

  void
  Clear();

  [[nodiscard]] bool
  Empty() const;

  void
  Append(Edge e);

  [[nodiscard]] std::span<const Edge>
  Export() const;

  private:
  std::array<Edge, Edge::Max> m;
  int len = 0;
};
