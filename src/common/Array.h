#pragma once

#include <initializer_list>

template <typename T, int Size>
class Array {
  public:
  T&
  At(int i) {
    return data[i];
  }

  const T&
  At(int i) const {
    return data[i];
  }

  Array&
  operator=(std::initializer_list<T> init) {
    int i = 0;
    for (const auto& item : init) {
      if (i < Size) {
        data[i++] = item;
      }
    }
    return *this;
  }

  auto
  begin() {
    return data;
  }

  auto
  end() {
    return data + Size;
  }

  auto
  begin() const {
    return data;
  }

  auto
  end() const {
    return data + Size;
  }

  private:
  T data[Size] = {};
};
