#pragma once

#include <QResizeEvent>
#include <QWidget>
#include <array>
#include <memory>

#include "../model/Edge.h"
#include "DotCanvas.h"
#include "EdgeCanvas.h"
#include "config.h"

class EdgeCanvasLayer final : public QWidget {
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