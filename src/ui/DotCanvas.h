#pragma once

#include <QPaintEvent>
#include <QPainter>
#include <QWidget>

#include "common.h"

class DotCanvas final : public QWidget {
  Q_OBJECT

  public:
  static constexpr int R = 8;
  static constexpr int A = 2 * R;

  static QColor
  Color();

  explicit DotCanvas(QWidget* parent = nullptr);

  void
  paintEvent(QPaintEvent* event) override;
};
