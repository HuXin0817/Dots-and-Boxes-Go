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
    Size = 1;
  }

  void
  Clear() {
    Size = 0;
  }

  [[nodiscard]] bool
  Empty() const {
    return Size == 0;
  }

  void
  Append(T e) {
    assert(len < Cap);
    Data.At(Size++) = e;
  }

  [[nodiscard]] Span<T>
  Export() {
    return {Data.begin(), Data.begin() + Size};
  }

  [[nodiscard]] Span<const T>
  Export() const {
    return {Data.begin(), Data.begin() + Size};
  }

  auto
  begin() {
    return Data.begin();
  }

  auto
  end() {
    return Data.begin() + Size;
  }

  auto
  begin() const {
    return Data.begin();
  }

  auto
  end() const {
    return Data.begin() + Size;
  }

  private:
  Array<T, Cap> Data;
  int Size = 0;
};
