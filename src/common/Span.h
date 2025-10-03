#pragma once

#include <random>

template <typename T>
class Span {
  public:
  Span(T* beg, T* end) : _begin(beg), _end(end) {
  }

  // Constructor for const conversion
  template <typename U>
  Span(const Span<U>& other) : _begin(other.begin()), _end(other.end()) {
  }

  auto
  begin() {
    return _begin;
  }

  auto
  end() {
    return _end;
  }

  auto
  begin() const {
    return _begin;
  }

  auto
  end() const {
    return _end;
  }

  auto
  Size() const {
    return _end - _begin;
  }

  auto
  size() const {
    return Size();
  }

  T&
  At(int i) {
    return _begin[i];
  }

  const T&
  At(int i) const {
    return _begin[i];
  }

  bool
  Empty() const {
    return _end == _begin;
  }

  private:
  T* _begin;
  T* _end;
};

template <typename T>
[[nodiscard]] const T&
RandomChoice(const Span<T>& data) {
  assert(!data.Empty());

  thread_local std::mt19937 rng(std::random_device{}());
  std::uniform_int_distribution<size_t> dist(0, data.Size() - 1);

  return data.At(dist(rng));
}
