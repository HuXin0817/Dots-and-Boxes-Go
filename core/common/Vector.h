#pragma once

template <class T>
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

  T*
  begin() {
    return Data;
  }

  T*
  end() {
    return Data + Len;
  }

  const T*
  begin() const {
    return Data;
  }

  const T*
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
