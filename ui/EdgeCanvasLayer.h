#pragma once

#include <QWidget>

#include "../core/common/Array.h"
#include "../core/model/Edge.h"
#include "BaseCanvasLayer.h"
#include "EdgeCanvas.h"
#include "EdgeLayer.h"

class EdgeCanvasLayer final : public EdgeLayer<EdgeCanvas> {
  friend class MainWindow;

  Q_OBJECT

  public:
  explicit EdgeCanvasLayer(QWidget* parent = nullptr) : EdgeLayer(parent) {
    resize(WindowSize, WindowSize);

    for (Edge edge = 0; edge < Edge::Max; edge++) {
      Canvases.At(edge).New(edge.Rotate(), this);
    }
  }
};
