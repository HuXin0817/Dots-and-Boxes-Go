#pragma once

#include "Common.h"
#include "DotCanvas.h"
#include "EdgeCanvas.h"

class BoxCanvas final : public QWidget {
  Q_OBJECT

  public:
  static constexpr int A = EdgeCanvas::B - 2 * DotCanvas::R;

  explicit BoxCanvas(QWidget* parent = nullptr) : QWidget(parent) {
    resize(QSize(A, A));
  }

  State state = State::Free;

  [[nodiscard]] QColor
  Color() const {
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

  protected:
  void
  paintEvent(QPaintEvent* event) override {
    QWidget::paintEvent(event);

    QPainter painter(this);
    painter.setRenderHint(QPainter::Antialiasing);
    painter.setPen(Qt::NoPen);
    painter.setBrush(QBrush(Color()));

    int x = width() / 2 - A / 2;
    int y = height() / 2 - A / 2;

    painter.drawRect(x, y, A, A);
  }
};
