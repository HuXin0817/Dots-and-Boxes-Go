#pragma once

#include <QPaintEvent>
#include <QPainter>
#include <QWidget>

#include "Common.h"

class DotCanvas final : public QWidget {
  Q_OBJECT

  public:
  static constexpr int R = 8;
  static constexpr int A = 2 * R;

  static QColor
  Color() {
    if (isDarkMode()) {
      return {202, 202, 202, 255};
    } else {
      return {255, 255, 255, 255};
    }
  }

  explicit DotCanvas(QWidget* parent = nullptr) : QWidget(parent) {
    setFixedSize(A, A);
  }

  void
  paintEvent(QPaintEvent* event) override {
    QWidget::paintEvent(event);

    QPainter painter(this);

    painter.setRenderHint(QPainter::Antialiasing);
    painter.setBrush(QBrush(Color()));
    painter.setPen(Qt::NoPen);

    int x = width() / 2;
    int y = height() / 2;

    painter.drawEllipse(QPoint(x, y), R, R);
  }
};
