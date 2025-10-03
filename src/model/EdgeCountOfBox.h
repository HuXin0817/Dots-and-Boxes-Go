#pragma once

#include "../common/Array.h"
#include "EdgeBoxMapper.h"

class EdgeCountOfBox : public Array<int, Box::Max> {
  public:
  EdgeCountOfBox() = default;

  int
  Add(Edge e);

  int
  MaxCount(Edge e) const;
};
