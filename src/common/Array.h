#pragma once

#include <initializer_list>

template <typename T, int Size>
class Array {
  public:
  T&
  At(int i) {
    return Data[i];
  }

  const T&
  At(int i) const {
    return Data[i];
  }

  Array&
  operator=(const Array& other) {
    if (this != &other) {
      std::memcpy(Data, other.Data, Size * sizeof(T));
    }
    return *this;
  }

  Array&
  operator=(std::initializer_list<T> init) {
    int i = 0;
    for (const auto& item : init) {
      if (i < Size) {
        Data[i++] = item;
      }
    }
    return *this;
  }

  auto
  begin() {
    return Data;
  }

  auto
  end() {
    return Data + Size;
  }

  auto
  begin() const {
    return Data;
  }

  auto
  end() const {
    return Data + Size;
  }

  private:
  T Data[Size] = {};
};
