#include "Box.h"

Box::Box(int v) : v(v) {
}

Box::Box(int x, int y) : v(x * BoardSize + y) {
}

Box::
operator int() const {
  return v;
}
