#pragma once

#include <cassert>

template <typename T, int Cap>
class Queue {
  public:
  void
  Clear() {
    front = 0;
    end = 0;
  }

  [[nodiscard]] bool
  Empty() const {
    return front == end;
  }

  void
  Append(T e) {
    assert(end < Cap);
    m.At(end++) = e;
  }

  T
  Pop() {
    assert(!Empty());
    auto e = m.At(front);
    front++;
    return e;
  }

  [[nodiscard]] Span<T>
  Export() const {
    return {m.begin() + front, m.begin() + end};
  }

  [[nodiscard]] bool
  Contains(T e) const {
    for (int i = front; i < end; i++) {
      if (static_cast<int>(m.At(i)) == static_cast<int>(e)) {
        return true;
      }
    }
    return false;
  }

  private:
  Array<T, Cap> m;
  int front = 0;
  int end = 0;
};
