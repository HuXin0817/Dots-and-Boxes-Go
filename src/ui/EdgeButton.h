#pragma once

#include <QPaintEvent>

#include "DotCanvas.h"
#include "common.h"

class EdgeButton final : public QWidget {
  Q_OBJECT

  public:
  static constexpr int A = DotCanvas::R * 2;
  static constexpr int B = A * 5;

  std::function<void()> CallBack;

  explicit EdgeButton(bool rotate, std::function<void()> CallBack, QWidget* parent = nullptr);

  bool Rotate = false;

  protected:
  void
  mousePressEvent(QMouseEvent* event) override;
};
