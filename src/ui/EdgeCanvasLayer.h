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
      if (Edge(i).dot1().X() == Edge(i).dot2().X()) {
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

    int x0 = (width() - BoardWindowSize) / 2 - DotCanvas::R;
    int y0 = (height() - BoardWindowSize) / 2 - DotCanvas::R;

    for (int i = 0; i < Edge::Max; i++) {
      Edge e(i);
      int x = x0 + e.dot1().X() * EdgeCanvas::B;
      int y = y0 + e.dot1().Y() * EdgeCanvas::B;
      if (e.dot1().X() == e.dot2().X()) {
        y += DotCanvas::R;
      } else {
        x += DotCanvas::R;
      }
      EdgeCanvases.At(e)->move(x, y);
    }
  }

  private:
  Array<std::unique_ptr<EdgeCanvas>, Edge::Max> EdgeCanvases;
};