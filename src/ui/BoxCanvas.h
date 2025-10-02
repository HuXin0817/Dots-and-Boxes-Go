#pragma once

#include "DotCanvas.h"
#include "EdgeCanvas.h"
#include "common.h"

class BoxCanvas final : public QWidget {
  Q_OBJECT

  public:
  static constexpr int A = EdgeCanvas::B - 2 * DotCanvas::R;

  explicit BoxCanvas(QWidget* parent = nullptr);

  State state = State::Free;

  [[nodiscard]] QColor
  Color() const;

  protected:
  void
  paintEvent(QPaintEvent* event) override;
};
