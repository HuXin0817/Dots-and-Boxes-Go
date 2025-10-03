#pragma once

#include <QWidget>
#include <memory>

#include "../common/Array.h"
#include "../model/Edge.h"
#include "BaseCanvasLayer.h"
#include "EdgeCanvas.h"
#include "EdgeLayer.h"

class EdgeCanvasLayer final : public EdgeLayer<EdgeCanvas> {
  friend class MainWindow;

  Q_OBJECT

  public:
  explicit EdgeCanvasLayer(QWidget* parent = nullptr) : EdgeLayer(parent) {
    resize(WindowSize, WindowSize);

    for (int edge = 0; edge < Edge::Max; edge++) {
      Canvases.At(edge) = std::make_unique<EdgeCanvas>(Edge(edge).Rotate(), this);
    }
  }
};
