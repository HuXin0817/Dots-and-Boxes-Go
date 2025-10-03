#pragma once

#include <cassert>

#include "Array.h"
#include "Span.h"

template <typename T, int Cap>
class Queue {
  public:
  void
  Clear() {
    Begin = 0;
    End = 0;
  }

  bool
  Empty() const {
    return Begin == End;
  }

  void
  Append(T e) {
    assert(End < Cap);
    Data.At(End++) = e;
  }

  T
  Pop() {
    assert(!Empty());
    auto e = Data.At(Begin);
    Begin++;
    return e;
  }

  Span<T>
  Export() const {
    return {Data.begin() + Begin, Data.begin() + End};
  }

  bool
  Contains(T e) const {
    for (int i = Begin; i < End; i++) {
      if (Data.At(i) == e) {
        return true;
      }
    }
    return false;
  }

  private:
  Array<T, Cap> Data;
  int Begin = 0;
  int End = 0;
};
