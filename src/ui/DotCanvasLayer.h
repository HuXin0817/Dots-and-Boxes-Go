#pragma once

#include <QResizeEvent>
#include <QWidget>
#include <array>
#include <memory>

#include "../model/Dot.h"
#include "DotCanvas.h"
#include "EdgeCanvas.h"
#include "config.h"

class DotCanvasLayer final : public QWidget {
  friend class MainWindow;

  Q_OBJECT

  public:
  explicit DotCanvasLayer(QWidget* parent = nullptr);

  protected:
  void
  resizeEvent(QResizeEvent* event) override;

  private:
  std::array<std::unique_ptr<DotCanvas>, Dot::Max> DotCanvases;
};