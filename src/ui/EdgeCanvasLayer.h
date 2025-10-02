#pragma once

#include <QResizeEvent>
#include <QWidget>
#include <array>
#include <memory>

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
  std::array<std::unique_ptr<EdgeCanvas>, Edge::Max> EdgeCanvases;
};