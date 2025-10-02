#pragma once

#include <QWidget>

#include "BoxCanvas.h"
#include "EdgeCanvas.h"

class BaseCanvasLayer : public QWidget {
  public:
  static constexpr int BoardWindowSize = Box::Size * EdgeCanvas::B;
  static constexpr int WindowSize = (Box::Size - 1) * EdgeCanvas::B + 2 * BoxCanvas::A + 80;

  explicit BaseCanvasLayer(QWidget* parent = nullptr) : QWidget(parent) {
  }
};