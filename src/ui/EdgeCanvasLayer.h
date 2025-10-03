#pragma once

#include <QWidget>
#include <memory>

#include "../common/Array.h"
#include "../model/Edge.h"
#include "BaseCanvasLayer.h"
#include "DotCanvas.h"
#include "EdgeCanvas.h"

class EdgeCanvasLayer final : public BaseCanvasLayer {
  friend class MainWindow;

  Q_OBJECT

  public:
  explicit EdgeCanvasLayer(QWidget* parent = nullptr) : BaseCanvasLayer(parent) {
    resize(WindowSize, WindowSize);

    for (int i = 0; i < Edge::Max; i++) {
      if (Edge(i).Dot1().X() == Edge(i).Dot2().X()) {
        EdgeCanvases.At(i) = std::make_unique<EdgeCanvas>(false, this);
      } else {
        EdgeCanvases.At(i) = std::make_unique<EdgeCanvas>(true, this);
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
      Edge e(i);
      int x = x0 + e.Dot1().X() * EdgeCanvas::B;
      int y = y0 + e.Dot1().Y() * EdgeCanvas::B;
      if (e.Dot1().X() == e.Dot2().X()) {
        y += R;
      } else {
        x += R;
      }
      EdgeCanvases.At(e)->move(x, y);
    }
  }

  private:
  Array<std::unique_ptr<EdgeCanvas>, Edge::Max> EdgeCanvases;
};