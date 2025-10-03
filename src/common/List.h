#pragma once

#include <cassert>

#include "Array.h"
#include "Span.h"

template <typename T, int Cap>
class List {
  public:
  void
  Reset(T e) {
    Data.At(0) = e;
    Len = 1;
  }

  void
  Clear() {
    Len = 0;
  }

  bool
  Empty() const {
    return Len == 0;
  }

  void
  Append(T e) {
    assert(Len < Cap);
    Data.At(Len++) = e;
  }

  Span<T>
  Export() const {
    return {Data.begin(), Data.begin() + Len};
  }

  int
  Size() const {
    return Len;
  }

  T*
  begin() {
    return Data.begin();
  }

  T*
  end() {
    return Data.begin() + Len;
  }

  const T*
  begin() const {
    return Data.begin();
  }

  const T*
  end() const {
    return Data.begin() + Len;
  }

  private:
  Array<T, Cap> Data;
  int Len = 0;
};
