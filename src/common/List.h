#pragma once

#include <cassert>

#include "Array.h"
#include "Span.h"

template <typename T, int Cap>
class List {
  public:
  void
  Reset(T e) {
    m.At(0) = e;
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
  Append(T e) {
    assert(len < Cap);
    m.At(len++) = e;
  }

  [[nodiscard]] Span<T>
  Export() {
    return {m.begin(), m.begin() + len};
  }

  [[nodiscard]] Span<const T>
  Export() const {
    return {m.begin(), m.begin() + len};
  }

  auto
  begin() {
    return m.begin();
  }

  auto
  end() {
    return m.begin() + len;
  }

  auto
  begin() const {
    return m.begin();
  }

  auto
  end() const {
    return m.begin() + len;
  }

  private:
  Array<T, Cap> m;
  int len = 0;
};
