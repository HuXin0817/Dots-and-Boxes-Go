#pragma once

#include <QWidget>
#include <memory>

#include "../common/Array.h"
#include "../model/Edge.h"
#include "EdgeButton.h"
#include "EdgeLayer.h"

class EdgeButtonLayer final : public EdgeLayer<EdgeButton> {
  friend class MainWindow;

  Q_OBJECT

  public:
  explicit EdgeButtonLayer(const std::function<std::function<void()>(Edge)>& CallBackFactory,
                           QWidget* parent = nullptr)
      : EdgeLayer(parent) {
    resize(WindowSize, WindowSize);

    for (int edge = 0; edge < Edge::Max; edge++) {
      if (Edge(edge).Dot1().X() == Edge(edge).Dot2().X()) {
        Canvases.At(edge) = std::make_unique<EdgeButton>(false, CallBackFactory(edge), this);
      } else {
        Canvases.At(edge) = std::make_unique<EdgeButton>(true, CallBackFactory(edge), this);
      }
    }
  }
};
