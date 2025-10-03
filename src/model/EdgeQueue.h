#pragma once

#include "Edge.h"

class EdgeQueue {
  public:
  void
  Clear();

  [[nodiscard]] bool
  Empty() const;

  void
  Append(Edge e);

  Edge
  Pop();

  [[nodiscard]] std::span<const Edge>
  Export() const;

  [[nodiscard]] bool
  Contains(Edge e) const;

  private:
  std::array<Edge, Edge::Max> m;
  int front = 0;
  int end = 0;
};
