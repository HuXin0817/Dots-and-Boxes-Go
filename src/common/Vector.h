#pragma once

template <typename T>
class Vector {
  public:
  explicit Vector(int size) : size(size) {
    data = new T[size];
  }

  ~Vector() {
    delete[] data;
  }

  T&
  At(int i) {
    return data[i];
  }

  const T&
  At(int i) const {
    return data[i];
  }

  auto
  begin() {
    return data;
  }

  auto
  end() {
    return data + size;
  }

  auto
  begin() const {
    return data;
  }

  auto
  end() const {
    return data + size;
  }

  int
  Size() const {
    return size;
  }

  private:
  T* data;
  int size;
};
