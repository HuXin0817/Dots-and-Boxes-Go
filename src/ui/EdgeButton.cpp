#include "EdgeButton.h"

EdgeButton::EdgeButton(bool rotate, std::function<void()> CallBack, QWidget* parent)
    : QWidget(parent), CallBack(std::move(CallBack)), Rotate(rotate) {
  if (!rotate) {
    resize(QSize(A, B));
  } else {
    resize(QSize(B, A));
  }
}

void
EdgeButton::mousePressEvent(QMouseEvent* event) {
  QWidget::mousePressEvent(event);

  CallBack();
}