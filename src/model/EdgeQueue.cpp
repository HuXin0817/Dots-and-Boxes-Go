#include "EdgeQueue.h"

void
EdgeQueue::Clear() {
  front = 0;
  end = 0;
}

[[nodiscard]] bool
EdgeQueue::Empty() const {
  return front == end;
}

void
EdgeQueue::Append(Edge e) {
  assert(end < Edge::Max);
  m[end++] = e;
}

Edge
EdgeQueue::Pop() {
  assert(!Empty());
  auto e = m[front];
  front++;
  return e;
}

[[nodiscard]] std::vector<Edge>
EdgeQueue::Export() const {
  return {m.begin() + front, m.begin() + end};
}

[[nodiscard]] bool
EdgeQueue::Contains(Edge e) const {
  for (int i = front; i < end; i++) {
    if (static_cast<int>(m[i]) == static_cast<int>(e)) {
      return true;
    }
  }
  return false;
}
