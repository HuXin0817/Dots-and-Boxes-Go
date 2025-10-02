#include "DotCanvasLayer.h"

DotCanvasLayer::DotCanvasLayer(QWidget* parent) : QWidget(parent) {
  resize(WindowSize, WindowSize);
  for (int i = 0; i < Dot::Max; i++) {
    DotCanvases[i] = std::make_unique<DotCanvas>(this);
  }
}

void
DotCanvasLayer::resizeEvent(QResizeEvent* event) {
  QWidget::resizeEvent(event);

  int x0 = (width() - BoardWindowSize) / 2 - DotCanvas::R;
  int y0 = (height() - BoardWindowSize) / 2 - DotCanvas::R;

  for (int i = 0; i < Dot::Size; i++) {
    for (int j = 0; j < Dot::Size; j++) {
      int x = x0 + i * EdgeCanvas::B;
      int y = y0 + j * EdgeCanvas::B;
      DotCanvases[Dot(i, j)]->move(x, y);
    }
  }
}
