#pragma once

#include <cassert>
#include <random>
#include <span>

static constexpr int BoardSize = 6;

template <typename T>
[[nodiscard]] const T&
RandomChoice(const std::span<T>& data) {
  assert(!data.empty());

  thread_local std::mt19937 rng(std::random_device{}());
  std::uniform_int_distribution<size_t> dist(0, data.size() - 1);

  return data[dist(rng)];
}

template <typename T, int Cap>
class List {
  public:
  void
  Reset(T e) {
    m[0] = e;
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
    m[len++] = e;
  }

  [[nodiscard]] std::span<const T>
  Export() const {
    return {m.begin(), m.begin() + len};
  }

  private:
  std::array<T, Cap> m;
  int len = 0;
};

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
    m[end++] = e;
  }

  T
  Pop() {
    assert(!Empty());
    auto e = m[front];
    front++;
    return e;
  }

  [[nodiscard]] std::span<const T>
  Export() const {
    return {m.begin() + front, m.begin() + end};
  }

  [[nodiscard]] bool
  Contains(T e) const {
    for (int i = front; i < end; i++) {
      if (m[i] == e) {
        return true;
      }
    }
    return false;
  }

  private:
  std::array<T, Cap> m;
  int front = 0;
  int end = 0;
};
