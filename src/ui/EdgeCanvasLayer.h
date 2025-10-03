#pragma once

#include <QWidget>
#include <memory>

#include "../common/Array.h"
#include "../model/Edge.h"
#include "BaseCanvasLayer.h"
#include "EdgeCanvas.h"

class EdgeCanvasLayer final : public BaseCanvasLayer {
  friend class MainWindow;

  Q_OBJECT

  public:
  explicit EdgeCanvasLayer(QWidget* parent = nullptr) : BaseCanvasLayer(parent) {
    resize(WindowSize, WindowSize);

    for (int edge = 0; edge < Edge::Max; edge++) {
      if (Edge(edge).Dot1().X() == Edge(edge).Dot2().X()) {
        EdgeCanvases.At(edge) = std::make_unique<EdgeCanvas>(false, this);
      } else {
        EdgeCanvases.At(edge) = std::make_unique<EdgeCanvas>(true, this);
      }
    }
  }

  protected:
  void
  resizeEvent(QResizeEvent* event) override {
    QWidget::resizeEvent(event);

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
      EdgeCanvases.At(edge)->move(x, y);
    }
  }

  private:
  Array<std::unique_ptr<EdgeCanvas>, Edge::Max> EdgeCanvases;
};
