#pragma once

#include <QPaintEvent>

#include "Common.h"
#include "DotCanvas.h"

class EdgeButton final : public QWidget {
  Q_OBJECT

  public:
  static constexpr int A = DotCanvas::R * 2;
  static constexpr int B = A * 5;

  std::function<void()> CallBack;

  explicit EdgeButton(bool rotate, std::function<void()> CallBack, QWidget* parent = nullptr)
      : QWidget(parent), CallBack(std::move(CallBack)), Rotate(rotate) {
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
    QWidget::mousePressEvent(event);

    CallBack();
  }
};
