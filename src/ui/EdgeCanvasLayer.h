#pragma once

#include <QResizeEvent>
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
  explicit EdgeCanvasLayer(QWidget* parent = nullptr);

  protected:
  void
  resizeEvent(QResizeEvent* event) override;

  private:
  Array<std::unique_ptr<EdgeCanvas>, Edge::Max> EdgeCanvases;
};