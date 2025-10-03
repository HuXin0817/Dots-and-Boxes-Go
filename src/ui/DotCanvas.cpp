#include "DotCanvas.h"

QColor
DotCanvas::Color() {
  if (isDarkMode()) {
    return {202, 202, 202, 255};
  } else {
    return {255, 255, 255, 255};
  }
}

DotCanvas::DotCanvas(QWidget* parent) : QWidget(parent) {
  setFixedSize(A, A);
}

void
DotCanvas::paintEvent(QPaintEvent* event) {
  QWidget::paintEvent(event);

  QPainter painter(this);

  painter.setRenderHint(QPainter::Antialiasing);
  painter.setBrush(QBrush(Color()));
  painter.setPen(Qt::NoPen);

  int x = width() / 2;
  int y = height() / 2;

  painter.drawEllipse(QPoint(x, y), R, R);
}
