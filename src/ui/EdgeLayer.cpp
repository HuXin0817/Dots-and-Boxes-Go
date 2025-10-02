#include "EdgeLayer.h"

EdgeLayer::EdgeLayer(QWidget* parent) : QWidget(parent) {
  resize(WindowSize, WindowSize);

  for (int i = 0; i < Edge::Max; i++) {
    if (Edge(i).dot1().X() == Edge(i).dot2().X()) {
      EdgeCanvases[i] = std::make_unique<EdgeCanvas>(false, this);
    } else {
      EdgeCanvases[i] = std::make_unique<EdgeCanvas>(true, this);
    }
  }
}

void
EdgeLayer::resizeEvent(QResizeEvent* event) {
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
    EdgeCanvases[e]->move(x, y);
  }
}
