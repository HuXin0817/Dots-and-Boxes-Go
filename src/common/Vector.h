#pragma once

template <typename T>
class Vector {
  public:
  explicit Vector(int Len) : Len(Len) {
    Data = new T[Len];
  }

  ~Vector() {
    delete[] Data;
  }

  T&
  At(int i) {
    return Data[i];
  }

  const T&
  At(int i) const {
    return Data[i];
  }

  auto
  begin() {
    return Data;
  }

  auto
  end() {
    return Data + Len;
  }

  auto
  begin() const {
    return Data;
  }

  auto
  end() const {
    return Data + Len;
  }

  int
  Size() const {
    return Len;
  }

  private:
  T* Data;
  int Len;
};
