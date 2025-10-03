#pragma once

#include <QWidget>
#include <memory>

#include "../common/Array.h"
#include "../model/Edge.h"
#include "BaseCanvasLayer.h"
#include "EdgeButton.h"

class EdgeButtonLayer final : public BaseCanvasLayer {
  friend class MainWindow;

  Q_OBJECT

  public:
  explicit EdgeButtonLayer(const std::function<std::function<void()>(Edge)>& CallBackFactory,
                           QWidget* parent = nullptr)
      : BaseCanvasLayer(parent) {
    resize(WindowSize, WindowSize);

    for (int edge = 0; edge < Edge::Max; edge++) {
      if (Edge(edge).Dot1().X() == Edge(edge).Dot2().X()) {
        EdgeButtons.At(edge) = std::make_unique<EdgeButton>(false, CallBackFactory(edge), this);
      } else {
        EdgeButtons.At(edge) = std::make_unique<EdgeButton>(true, CallBackFactory(edge), this);
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
      int x = x0 + edge.Dot1().X() * EdgeButton::B;
      int y = y0 + edge.Dot1().Y() * EdgeButton::B;
      if (edge.Dot1().X() == edge.Dot2().X()) {
        y += R;
      } else {
        x += R;
      }
      EdgeButtons.At(edge)->move(x, y);
    }
  }

  private:
  Array<std::unique_ptr<EdgeButton>, Edge::Max> EdgeButtons;
};
