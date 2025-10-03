#pragma once

#include <random>

template <typename T>
class Span {
  public:
  Span() = default;

  Span(const T* begin, const T* end) : Begin(begin), End(end) {
  }

  const T*
  begin() const {
    return Begin;
  }

  const T*
  end() const {
    return End;
  }

  int
  Size() const {
    return End - Begin;
  }

  const T&
  At(int i) const {
    return Begin[i];
  }

  bool
  Empty() const {
    return End == Begin;
  }

  private:
  const T* Begin = nullptr;
  const T* End = nullptr;
};

template <typename T>
const T&
RandomChoice(const Span<T>& data) {
  assert(!data.Empty());

  thread_local std::mt19937 rng(std::random_device{}());
  std::uniform_int_distribution<size_t> dist(0, data.Size() - 1);

  return data.At(dist(rng));
}
