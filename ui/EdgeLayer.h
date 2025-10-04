#pragma once

#include "../core/common/Array.h"
#include "../core/model/Edge.h"
#include "BaseCanvasLayer.h"

template <class Canvas>
class EdgeLayer : public BaseCanvasLayer {
  friend class MainWindow;

  public:
  explicit EdgeLayer(QWidget* parent = nullptr) : BaseCanvasLayer(parent) {
  }

  protected:
  void
  resizeEvent(QResizeEvent* event) override {
    BaseCanvasLayer::resizeEvent(event);

    int x0 = (width() - BoardWindowSize) / 2 - R;
    int y0 = (height() - BoardWindowSize) / 2 - R;

    for (int i = 0; i < Edge::Max; i++) {
      Edge edge(i);
      int x = x0 + edge.Dot1().X() * EdgeCanvas::B;
      int y = y0 + edge.Dot1().Y() * EdgeCanvas::B;
      if (edge.Dot1().X() == edge.Dot2().X()) {
        y += R;
      } else {
        x += R;
      }
      Canvases.At(edge)->move(x, y);
    }
  }

  Array<Ptr<Canvas>, Edge::Max> Canvases;
};
