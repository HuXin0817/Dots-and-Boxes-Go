#pragma once

#include "DotCanvas.h"

class EdgeButton final : public BaseCanvas {
  Q_OBJECT

  public:
  static constexpr int A = R * 2;
  static constexpr int B = A * 5;

  std::function<void()> CallBack;

  explicit EdgeButton(bool rotate, std::function<void()> CallBack, QWidget* parent = nullptr)
      : BaseCanvas(parent), CallBack(std::move(CallBack)), Rotate(rotate) {
    if (!rotate) {
      resize(QSize(A, B));
    } else {
      resize(QSize(B, A));
    }
  }

  bool Rotate = false;

  protected:
  void
  mousePressEvent(QMouseEvent* event) override {
    BaseCanvas::mousePressEvent(event);

    CallBack();
  }
};
