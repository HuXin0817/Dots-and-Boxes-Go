#include "BoxCanvas.h"

BoxCanvas::BoxCanvas(QWidget* parent) : QWidget(parent) {
  resize(QSize(A, A));
}

QColor
BoxCanvas::Color() const {
  if (state == State::Free) {
    return {0, 0, 0, 0};
  }
  if (state == State::Player1Occupy) {
    return {64, 64, 255, 64};
  }
  if (state == State::Player2Occupy) {
    return {255, 64, 64, 64};
  }

  return {};
}

void
BoxCanvas::paintEvent(QPaintEvent* event) {
  QWidget::paintEvent(event);

  QPainter painter(this);
  painter.setRenderHint(QPainter::Antialiasing);
  painter.setPen(Qt::NoPen);
  painter.setBrush(QBrush(Color()));

  int x = width() / 2 - A / 2;
  int y = height() / 2 - A / 2;

  painter.drawRect(x, y, A, A);
}
