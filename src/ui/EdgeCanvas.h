#pragma once

#include <QPaintEvent>

#include "Common.h"
#include "DotCanvas.h"

class EdgeCanvas final : public QWidget {
  Q_OBJECT

  public:
  static constexpr int A = DotCanvas::R * 2;
  static constexpr int B = A * 5;

  explicit EdgeCanvas(bool rotate, QWidget* parent = nullptr) : QWidget(parent), Rotate(rotate) {
    if (!rotate) {
      resize(QSize(A, B));
    } else {
      resize(QSize(B, A));
    }
  }

  State state = State::Free;
  bool highLight = true;
  bool Rotate = false;

  QColor
  Color() const {
    if (state == State::Free) {
      if (isDarkMode()) {
        return {65, 65, 65, 255};
      }
      return {217, 217, 217, 255};
    }

    QColor color;
    if (state == State::Player1Occupy) {
      color = {64, 64, 255, 255};
    }

    if (state == State::Player2Occupy) {
      color = {255, 64, 64, 255};
    }

    if (highLight) {
      color.setAlpha(255);
    } else {
      color.setAlpha(128);
    }

    return color;
  }

  protected:
  void
  paintEvent(QPaintEvent* event) override {
    QWidget::paintEvent(event);

    QPainter painter(this);
    painter.setRenderHint(QPainter::Antialiasing);
    painter.setPen(Qt::NoPen);
    painter.setBrush(QBrush(Color()));

    if (!Rotate) {
      int x = width() / 2 - A / 2;
      int y = height() / 2 - B / 2;
      painter.drawRect(x, y, A, B);
    } else {
      int x = width() / 2 - B / 2;
      int y = height() / 2 - A / 2;
      painter.drawRect(x, y, B, A);
    }
  }
};
