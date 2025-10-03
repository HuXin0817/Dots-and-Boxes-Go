#include "EdgeButtonLayer.h"

EdgeButtonLayer::EdgeButtonLayer(const std::function<std::function<void()>(Edge)>& CallBackFactory,
                                 QWidget* parent)
    : BaseCanvasLayer(parent) {
  resize(WindowSize, WindowSize);

  for (int i = 0; i < Edge::Max; i++) {
    if (Edge(i).dot1().X() == Edge(i).dot2().X()) {
      EdgeButtons.At(i) = std::make_unique<EdgeButton>(false, CallBackFactory(i), this);
    } else {
      EdgeButtons.At(i) = std::make_unique<EdgeButton>(true, CallBackFactory(i), this);
    }
  }
}

void
EdgeButtonLayer::resizeEvent(QResizeEvent* event) {
  QWidget::resizeEvent(event);

  int x0 = (width() - BoardWindowSize) / 2 - DotCanvas::R;
  int y0 = (height() - BoardWindowSize) / 2 - DotCanvas::R;

  for (int i = 0; i < Edge::Max; i++) {
    Edge e(i);
    int x = x0 + e.dot1().X() * EdgeButton::B;
    int y = y0 + e.dot1().Y() * EdgeButton::B;
    if (e.dot1().X() == e.dot2().X()) {
      y += DotCanvas::R;
    } else {
      x += DotCanvas::R;
    }
    EdgeButtons.At(e)->move(x, y);
  }
}
