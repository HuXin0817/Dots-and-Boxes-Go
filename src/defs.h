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