#pragma once

#include <QWidget>
#include <memory>

#include "../common/Array.h"
#include "../model/Box.h"
#include "BaseCanvasLayer.h"
#include "BoxCanvas.h"
#include "DotCanvas.h"
#include "EdgeCanvas.h"

class BoxCanvasLayer final : public BaseCanvasLayer {
  friend class MainWindow;

  Q_OBJECT

  public:
  explicit BoxCanvasLayer(QWidget* parent = nullptr) : BaseCanvasLayer(parent) {
    resize(WindowSize, WindowSize);
    for (int i = 0; i < Box::Max; i++) {
      BoxCanvases.At(i) = std::make_unique<BoxCanvas>(this);
    }
  }

  protected:
  void
  resizeEvent(QResizeEvent* event) override {
    QWidget::resizeEvent(event);

    int x0 = (width() - BoardWindowSize) / 2 + DotCanvas::R;
    int y0 = (height() - BoardWindowSize) / 2 + DotCanvas::R;

    for (int i = 0; i < Box::Size; i++) {
      for (int j = 0; j < Box::Size; j++) {
        int x = x0 + i * EdgeCanvas::B;
        int y = y0 + j * EdgeCanvas::B;
        BoxCanvases.At(Box(i, j))->move(x, y);
      }
    }
  }

  private:
  Array<std::unique_ptr<BoxCanvas>, Box::Max> BoxCanvases;
};