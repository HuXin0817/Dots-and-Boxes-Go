#pragma once

#include <array>

#include "EdgeBoxMapper.h"

class EdgeCountOfBox : public std::array<int, Box::Max> {
  public:
  EdgeCountOfBox() = default;

  int
  Add(Edge e);

  int
  MaxCount(Edge e) const;
};
