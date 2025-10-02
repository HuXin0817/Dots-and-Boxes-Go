#pragma once

#include <QPaintEvent>

#include "DotCanvas.h"
#include "common.h"

class EdgeCanvas final : public QWidget {
  Q_OBJECT

  public:
  static constexpr int A = DotCanvas::R * 2;
  static constexpr int B = A * 5;

  explicit EdgeCanvas(bool rotate, QWidget* parent = nullptr);

  State state = State::Free;
  bool highLight = true;
  bool Rotate = false;

  [[nodiscard]] QColor
  Color() const;

  protected:
  void
  paintEvent(QPaintEvent* event) override;
};
