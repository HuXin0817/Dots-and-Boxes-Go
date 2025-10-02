#include "DotCanvas.h"

QColor
DotCanvas::Color() {
  if (isDarkMode()) {
    return DarkThemeColor;
  } else {
    return LightThemeColor;
  }
}

DotCanvas::DotCanvas(QWidget* parent) : QWidget(parent) {
  setFixedSize(A, A);
}

void
DotCanvas::paintEvent(QPaintEvent* event) {
  QWidget::paintEvent(event);

  QPainter painter(this);

  painter.setRenderHint(QPainter::Antialiasing, true);
  painter.setBrush(QBrush(Color()));
  painter.setPen(Qt::NoPen);

  int x = width() / 2;
  int y = height() / 2;

  painter.drawEllipse(QPoint(x, y), R, R);
}
