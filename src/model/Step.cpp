#include "Step.h"

[[nodiscard]] bool
Step::Gaming() const {
  return v < Edge::Max;
}

[[nodiscard]] int
Step::RemainStep() const {
  return Edge::Max - v;
}

[[nodiscard]] int
Step::NowStep() const {
  return v;
}

void
Step::Go() {
  v++;
}
