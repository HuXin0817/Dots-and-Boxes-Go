#include "Dot.h"

Dot::Dot(int v) : v(v) {
}

Dot::Dot(int x, int y) : v(x * Size + y) {
}

[[nodiscard]] int
Dot::X() const {
  return v / Size;
}

[[nodiscard]] int
Dot::Y() const {
  return v % Size;
}

Dot::
operator int() const {
  return v;
}