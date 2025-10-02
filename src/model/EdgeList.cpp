#include "EdgeList.h"

void
EdgeList::Reset(Edge e) {
  m[0] = e;
  len = 1;
}

void
EdgeList::Clear() {
  len = 0;
}

[[nodiscard]] bool
EdgeList::Empty() const {
  return len == 0;
}

void
EdgeList::Append(Edge e) {
  assert(len < Edge::Max);
  m[len++] = e;
}

[[nodiscard]] std::span<const Edge>
EdgeList::Export() const {
  return {m.begin(), m.begin() + len};
}
