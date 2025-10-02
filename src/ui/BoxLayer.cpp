#include "BoxLayer.h"

BoxLayer::BoxLayer(QWidget* parent) : QWidget(parent) {
  resize(WindowSize, WindowSize);
  for (int i = 0; i < Box::Max; i++) {
    BoxCanvases[i] = std::make_unique<BoxCanvas>(this);
  }
}

void
BoxLayer::resizeEvent(QResizeEvent* event) {
  QWidget::resizeEvent(event);

  int x0 = (width() - BoardWindowSize) / 2 + DotCanvas::R;
  int y0 = (height() - BoardWindowSize) / 2 + DotCanvas::R;

  for (int i = 0; i < Box::Size; i++) {
    for (int j = 0; j < Box::Size; j++) {
      int x = x0 + i * EdgeCanvas::B;
      int y = y0 + j * EdgeCanvas::B;
      BoxCanvases[Box(i, j)]->move(x, y);
    }
  }
}
