#pragma once

#include <QWidget>

#include "../core/common/Array.h"
#include "../core/model/Dot.h"
#include "BaseCanvasLayer.h"
#include "DotCanvas.h"
#include "EdgeCanvas.h"

class DotCanvasLayer final : public BaseCanvasLayer {
  friend class MainWindow;

  Q_OBJECT

  public:
  explicit DotCanvasLayer(QWidget* parent = nullptr) : BaseCanvasLayer(parent) {
    resize(WindowSize, WindowSize);
    for (int dot = 0; dot < Dot::Max; dot++) {
      DotCanvases.At(dot).New(this);
    }
  }

  protected:
  void
  resizeEvent(QResizeEvent* event) override {
    BaseCanvasLayer::resizeEvent(event);

    int x0 = (width() - BoardWindowSize) / 2 - R;
    int y0 = (height() - BoardWindowSize) / 2 - R;

    for (int i = 0; i < Dot::Size; i++) {
      for (int j = 0; j < Dot::Size; j++) {
        int x = x0 + i * EdgeCanvas::B;
        int y = y0 + j * EdgeCanvas::B;
        DotCanvases.At(Dot(i, j))->move(x, y);
      }
    }
  }

  private:
  Array<Ptr<DotCanvas>, Dot::Max> DotCanvases;
};
